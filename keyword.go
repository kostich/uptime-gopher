package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func keywordHost(host, keyword string, ch chan<- string) {
	// check for protocol, if not defined use http
	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "http://" + host
	}

	keywordFound := false
	resp, err := http.Get(host)
	if err != nil {
		ch <- fmt.Sprintf("host: %v, keyword: %v, state: error: %v", host, keyword, err)
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

	ch <- fmt.Sprintf("host: %v, keyword: %v, state: %v", host, keyword, keywordFound)
}
