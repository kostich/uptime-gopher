package main

import (
	"fmt"
	"os/exec"
	"strings"
)

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
func pingHost(host string, ch chan<- string) {
	// first check if we have nmap on the system
	if !pingAvailable() {
		ch <- fmt.Sprintf("host: %v, state: ping tool not available", host)
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

	out, err := exec.Command("ping", host, "-c 4", "-w 10").Output()
	if err != nil {
		ch <- fmt.Sprintf("host: %v, state: error checking host: %v", host, err)
		return
	}
	if strings.Contains(string(out), "Destination Host Unreachable") {
		ch <- fmt.Sprintf("host: %v, state: offline", host)
	} else {
		ch <- fmt.Sprintf("host: %v, state: online", host)
	}
}
