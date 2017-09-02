// Uptime Gopher, a service for checking uptimes of various resources
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type host struct {
	Name         string   `json:"name"`
	Protocol     string   `json:"protocol"`
	Keywords     []string `json:"keywords"`
	Ping         bool     `json:"ping"`
	HTTP         bool     `json:"http"`
	HTTPResponse int      `json:"httpResponse"`
	Ports        []int    `json:"ports"`
}

// Check if all the requirements are satisfied.
func checkReqs() error {
	if _, err := os.Stat("hosts.json"); os.IsNotExist(err) {
		return fmt.Errorf("hosts.json file doesn't exist")
	}

	return nil
}

// Returns the full domain name, without protocol and path.
func justDomain(url string) string {
	domain := ""
	if strings.HasPrefix(url, "https://") {
		url = strings.TrimPrefix(url, "https://")
	} else if strings.HasPrefix(url, "http://") {
		url = strings.TrimPrefix(url, "http://")
	}

	for _, v := range url {
		if string(v) != "/" {
			domain += string(v)
		} else {
			return domain
		}
	}

	return domain
}

func main() {
	// We check the requirements
	err := checkReqs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "uptime-gopher: requirements not satisfied: %v\n", err)
		os.Exit(1)
	}
	// We read the hosts config file
	hostConf, err := ioutil.ReadFile("./hosts.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "uptime-gopher: error reading hosts config: %v\n", err)
		os.Exit(1)
	}
	var hosts []host
	json.Unmarshal(hostConf, &hosts)

	// check web capabilities, ping, ports and keywords
	webget := make(chan string)
	ping := make(chan string)
	ports := make(chan string)
	keywords := make(chan string)
	for _, h := range hosts {
		if h.Ping {
			go pingHost(h.Name, ping)
		}

		if h.HTTP {
			name := h.Protocol + "://" + h.Name
			go fetchHost(name, h.HTTPResponse, webget)
		}

		if len(h.Ports) != 0 {
			for _, p := range h.Ports {
				go portCheckHost(h.Name, p, ports)
			}
		}

		if len(h.Keywords) != 0 {
			for _, k := range h.Keywords {
				go keywordHost(h.Name, k, keywords)
			}
		}
	}

	// Output the results
	for _, h := range hosts {
		if h.HTTP {
			fmt.Printf("[WEBGET] %s\n", <-webget)
		}
	}

	for _, h := range hosts {
		if h.Ping {
			fmt.Printf("[ PING ] %s\n", <-ping)
		}
	}

	for _, h := range hosts {
		if len(h.Ports) != 0 {
			for i := len(h.Ports); i > 0; i-- {
				fmt.Printf("[ PORT ] %s\n", <-ports)
			}
		}
	}

	for _, h := range hosts {
		if len(h.Keywords) != 0 {
			for i := len(h.Keywords); i > 0; i-- {
				fmt.Printf("[KEYWRD] %s\n", <-keywords)
			}
		}
	}
}
