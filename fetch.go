package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Fetch a given host and return a status based on the response
func fetchHost(host string, response int, ch chan<- string) {
	// check for protocol, if not defined, use http
	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "http://" + host
	}

	resp, err := http.Get(host)
	if err != nil {
		ch <- fmt.Sprintf("host: %v, state: error: %v", host, err)
		return
	}
	defer resp.Body.Close()
	status := "OK"
	if resp.StatusCode != response {
		status = "NOT OK"
	}

	ch <- fmt.Sprintf("host: %v, response: %v, state: %v", host, resp.StatusCode, status)
}
