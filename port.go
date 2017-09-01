package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/net/idna"
)

func nmapAvailable() bool {
	cmd := exec.Command("which", "nmap")
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		return false
	}
	fmt.Println(err)
	return true
}

// Check if the given port is open on a given host, via nmap(1)
func portCheckHost(host string, port int, ch chan<- string) {
	// first check if we have nmap on the system
	if !nmapAvailable() {
		ch <- fmt.Sprintf("host: %v, port: %v, state: nmap tool not available", host, port)
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
		ch <- fmt.Sprintf("host: %v, port: %v, state: error determining punnycode for host: %v", host, port, err)
		return
	}
	out, err := exec.Command("nmap", "-sT", "-p", strPort, punnyHost).Output()
	if err != nil {
		ch <- fmt.Sprintf("host: %v, port: %v, state: error nmaping: %v", host, port, err)
		return
	}

	if strings.Contains(string(out), "open") {
		ch <- fmt.Sprintf("host: %v, port: %v, state: OPEN", host, port)
	} else if strings.Contains(string(out), "closed") {
		ch <- fmt.Sprintf("host: %v, port: %v, state: CLOSED", host, port)
	} else if strings.Contains(string(out), "filtered") {
		ch <- fmt.Sprintf("host: %v, port: %v, state: FILTERED", host, port)
	} else if strings.Contains(string(out), "resolve") {
		ch <- fmt.Sprintf("host: %v, state: UNRESOLVABLE", host)
	} else {
		ch <- fmt.Sprintf("host: %v, port: %v, state: UNKNOWN", host, port)
	}
}
