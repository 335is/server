package main

/*
	This program exercises the server REST API for testing, pushing metrics, etc.
*/

import (
	"fmt"
	"net/http"
	"time"
)

const (
	address    string        = "http://localhost"
	port       string        = ":2000"
	delay      time.Duration = 10 * time.Millisecond
	apiKeyName               = "APIKey"
	apiKey     string        = "SERVER-APIKEY-473b29ba-4ab3-46fa-bda1-9015444d70b5"
)

var ()

func main() {
	base := address + port
	paths := getPaths()
	for {
		for _, path := range paths {
			makeCall(base+path, delay)
			time.Sleep(delay)
		}
	}
}

func makeCall(path string, delay time.Duration) {
	req, err := http.NewRequest("GET", path, nil)
	req.Header.Add(apiKeyName, apiKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c := http.Client{}
	start := time.Now()
	resp, err := c.Do(req)
	duration := time.Since(start)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		fmt.Printf("HTTP status %d, %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
		return
	}

	fmt.Printf("%10dms: %s\n", duration.Milliseconds(), path)

	c.CloseIdleConnections()
}

// sneaky way to get a const slice of strings
func getPaths() []string {
	return []string{
		"/",
		"/favicon.ico",
		"/metrics",
		"/bands",
		"/bands/names",
		"/bands/beatles",
		"/bands/beatles/name",
		"/bands/beatles/members",
		"/bands/beatles/members/John%20Lennon",
		"/bands/beatles/members/John%20Lennon/name",
		"/bands/beatles/members/John%20Lennon/instruments",
		"/bands/beatles/members/John%20Lennon/instruments/vocals",
		"/bands/beatles/members/John%20Lennon/founder",
		"/bands/beatles/members/John%20Lennon/current",
		"/bands/beatles/year",
	}
}
