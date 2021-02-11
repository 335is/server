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
	address string        = "http://localhost"
	port    string        = ":2000"
	delay   time.Duration = 10 * time.Millisecond
)

var ()

func main() {
	base := address + port
	paths := getPaths()

	for {
		makeCalls(base, paths, delay)
	}
}

func makeCalls(base string, paths []string, delay time.Duration) {
	for _, path := range paths {
		p := base + path
		req, err := http.NewRequest("GET", p, nil)
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

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			fmt.Printf("HTTP status %d, %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
			return
		}

		fmt.Printf("%10dms: %s\n", duration.Milliseconds(), p)

		time.Sleep(delay)
	}
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
