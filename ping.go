package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/net/idna"
)

type pingResp struct {
	datetime time.Time
	host     string
	state    bool
	comment  string
}

// Checks if ping(1) tool is available.
func pingAvailable() bool {
	cmd := exec.Command("which", "ping")
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		return false
	}
	return true
}

// Ping the given host via ping(1).
func pingHost(host string, ch chan<- *pingResp) {
	// first check if we have nmap on the system
	if !pingAvailable() {
		r := pingResp{time.Now(), host, false, "ping tool not available"}
		ch <- &r
		return
	}

	// we don't need the protocol to ping the host
	if strings.HasPrefix(host, "https://") {
		host = strings.TrimPrefix(host, "https://")
	} else if strings.HasPrefix(host, "http://") {
		host = strings.TrimPrefix(host, "http://")
	}

	// we just need the domain
	host = justDomain(host)

	// transform domain to punnycode
	punnyStruct := idna.New()
	punnyHost, err := punnyStruct.ToASCII(host)
	if err != nil {
		r := pingResp{time.Now(), host, false, fmt.Sprintf("can't determine punnycode for host, %v", err)}
		ch <- &r
		return
	}

	out, err := exec.Command("ping", punnyHost, "-c4", "-w10").CombinedOutput()
	if err != nil {
		r := pingResp{time.Now(), host, false, "error checking host: " + string(out)}
		ch <- &r
		return
	}
	if strings.Contains(string(out), "Destination Host Unreachable") {
		r := pingResp{time.Now(), host, false, ""}
		ch <- &r
	} else {
		r := pingResp{time.Now(), host, true, ""}
		ch <- &r
	}
}
