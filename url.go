package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
)

//FetchURL fetches a http URL
func FetchURL(url string) []byte {

	if !IsURL(url) {
		app.Fatalf("URL is not valid: %s", url)
	}

	resp, err := http.Get(url)

	app.FatalIfError(err, "Failed to connect to url")

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	app.FatalIfError(err, "Failed to read response body")

	return body
}

//IsURL checks if a URL is valid
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func PostDashboard(host string, token string, dashboard *Dashboard) {

	if !IsURL(host) {
		app.Fatalf("Host URL is not valid: %s", host)
	}

	GD := struct {
		Dashboard *Dashboard `json:"dashboard"`
		Version   int        `json:"version"`
		Overwrite bool       `json:"overwrite"`
	}{
		Dashboard: dashboard,
		Version:   1,
		Overwrite: true,
	}

	body, err := json.Marshal(GD)

	app.FatalIfError(err, "Could not create json body")
	b := bytes.NewReader(body)
	req, err := http.NewRequest("POST", host+"/api/dashboards/db", b)

	app.FatalIfError(err, "Failed to create new request")

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	app.FatalIfError(err, "Failed to post dashboard")
	log.Info().Msgf("Request Succesful - Dashboard [%s] created!", dashboard.Title)
	defer resp.Body.Close()
}
