package main

import (
	"io/ioutil"
	"net/http"
	"net/url"

	log "k8s.io/klog"
)

//FetchURL fetches a http URL
func FetchURL(url string) []byte {
	if !IsURL(url) {
		log.Fatalf("URL is not valid: %s", url)
	}

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Failed to connect to url, ", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Failed to read response body, ", err.Error())
	}
	return body
}

//IsURL checks if a URL is valid
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
