package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type fetchResp struct {
	datetime    time.Time
	host        string
	desiredResp int
	actualResp  int
	comment     string
}

// Fetch a given host and return a status based on the response.
func fetchHost(host string, response int, ch chan<- *fetchResp) {
	// check for protocol, if not defined, use http
	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "http://" + host
	}

	resp, err := http.Get(host)
	if err != nil {
		r := fetchResp{time.Now(), host, response, 0, fmt.Sprintf("error, %v", err)}
		ch <- &r
		return
	}
	defer resp.Body.Close()

	r := fetchResp{time.Now(), host, response, resp.StatusCode, ""}
	ch <- &r
}
