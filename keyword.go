package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type keywordResp struct {
	datetime time.Time
	host     string
	keyword  string
	state    bool
	comment  string
}

// Checks if a HTML page on a given post contains <meta name="keywords"> with
// content which contains keyword.
func keywordHost(host, keyword string, ch chan<- *keywordResp) {
	// check for protocol, if not defined use http
	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "http://" + host
	}

	keywordFound := false
	resp, err := http.Get(host)
	if err != nil {
		r := keywordResp{time.Now(), host, keyword, keywordFound, fmt.Sprintf("%v", err)}
		ch <- &r
		return
	}
	defer resp.Body.Close()

	tzr := html.NewTokenizer(resp.Body)

	for tt := tzr.Next(); tt != html.ErrorToken; {
		token := tzr.Token()
		if token.Data == "meta" {
			for i, a := range token.Attr {
				if a.Val == "keywords" {
					// the next attribute contains the keywords then
					if strings.Contains(token.Attr[i+1].Val, keyword) {
						keywordFound = true
					}
				}
			}
		}
		tt = tzr.Next()
	}

	r := keywordResp{time.Now(), host, keyword, keywordFound, ""}
	ch <- &r
}
