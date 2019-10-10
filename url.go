package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//FetchURL fetches a http URL
func FetchURL(url string) []byte {
	if !IsURL(url) {
		log.Fatalf("URL is not valid: %s", url)
	}

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Failed to connect to %s Error:[%v]", url, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Failed to read response body Error:[%v]", err)
	}
	return body
}

//IsURL checks if a URL is valid
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
