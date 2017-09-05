package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/idna"
)

type portResp struct {
	datetime time.Time
	host     string
	port     int
	comment  string
}

// Checks if nmap(1) tool is available.
func nmapAvailable() bool {
	cmd := exec.Command("which", "nmap")
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		return false
	}
	return true
}

// Check if the given port is open on a given host, via nmap(1).
func portCheckHost(host string, port int, ch chan<- *portResp) {
	// first check if we have nmap on the system
	if !nmapAvailable() {
		r := portResp{time.Now(), host, port, "nmap tool not available"}
		ch <- &r
		return
	}

	// trim the protocol from the host, if any
	if strings.HasPrefix(host, "https://") {
		host = strings.TrimPrefix(host, "https://")
	} else if strings.HasPrefix(host, "http://") {
		host = strings.TrimPrefix(host, "http://")
	}

	// we just need the domain
	host = justDomain(host)
	strPort := strconv.Itoa(port)

	punnyStruct := idna.New()
	punnyHost, err := punnyStruct.ToASCII(host)
	if err != nil {
		r := portResp{time.Now(), host, port, fmt.Sprintf("can't determine punnycode for host, %v", err)}
		ch <- &r
		return
	}
	out, err := exec.Command("nmap", "-sT", "-p", strPort, punnyHost).CombinedOutput()
	if err != nil {
		r := portResp{time.Now(), host, port, fmt.Sprintf("error nmaping, %v, %v", string(out), err)}
		ch <- &r
		return
	}

	if strings.Contains(string(out), "open") {
		r := portResp{time.Now(), host, port, "OPEN"}
		ch <- &r
	} else if strings.Contains(string(out), "closed") {
		r := portResp{time.Now(), host, port, "CLOSED"}
		ch <- &r
	} else if strings.Contains(string(out), "filtered") {
		r := portResp{time.Now(), host, port, "FILTERED"}
		ch <- &r
	} else if strings.Contains(string(out), "resolve") {
		r := portResp{time.Now(), host, port, "UNRESOLVABLE"}
		ch <- &r
	} else {
		r := portResp{time.Now(), host, port, fmt.Sprintf("UNKNOWN, %v", string(out))}
		ch <- &r
	}
}
