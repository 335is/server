package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/335is/log"
	"github.com/335is/server/internal/data"
	"github.com/335is/server/internal/metrics"
	"github.com/335is/server/internal/middleware"
	"github.com/gorilla/mux"
)

// Example of a REST API that serves bands, members, instruments, etc.

var (
	content string = ""
)

// ServeHTTP is a blocking call the begins the web server
func ServeHTTP(port string, contentDir string) {
	content = contentDir

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Root)
	r.HandleFunc("/favicon.ico", FavIconHandler)
	r.HandleFunc("/metrics", MetricsHandler)
	r.HandleFunc("/bands", Bands)
	r.HandleFunc("/bands/names", BandNames)
	r.HandleFunc("/bands/{bandID}", Band)
	r.HandleFunc("/bands/{bandID}/name", BandName)
	r.HandleFunc("/bands/{bandID}/members", BandMembers)
	r.HandleFunc("/bands/{bandID}/members/{memberID}", BandMember)
	r.HandleFunc("/bands/{bandID}/members/{memberID}/name", BandMemberName)
	r.HandleFunc("/bands/{bandID}/members/{memberID}/instruments", BandMemberInstruments)
	r.HandleFunc("/bands/{bandID}/members/{memberID}/instruments/{instrumentID}", BandMemberInstrument)
	r.HandleFunc("/bands/{bandID}/members/{memberID}/founder", BandMemberFounder)
	r.HandleFunc("/bands/{bandID}/members/{memberID}/current", BandMemberCurrent)
	r.HandleFunc("/bands/{bandID}/year", BandYear)

	p := ":" + port
	log.Infof("HTTP server listening on %s", p)

	http.ListenAndServe(p, middleware.MetricsMiddleware(middleware.PanicMiddleware(middleware.LoggingMiddleware(r))))
}

// API route handlers

// Root handles requests to the root page
// http://localhost:80/
func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Welcome to the band server.")
}

// FavIconHandler - server up the browser icon
func FavIconHandler(w http.ResponseWriter, r *http.Request) {
	path, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	p := filepath.Join(path, content, "favicon.ico")
	http.ServeFile(w, r, p)
}

// MetricsHandler - displays accumulated metrics
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	stats := metrics.CalculateStatistics()
	serialize(stats, w)
}

// Bands returns the list of bands
// http://localhost:80/bands
func Bands(w http.ResponseWriter, r *http.Request) {
	b, err := data.GetBands()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(b, w)
}

// BandNames returns the list of band names
// http://localhost:80/bands/names
func BandNames(w http.ResponseWriter, r *http.Request) {
	names, err := data.GetBandNames()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(names, w)
}

// Band returns information about the specified band
// http://localhost:80/bands/bandID
func Band(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bandID"]

	band, err := data.GetBand(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(band, w)
}

// BandName returns the name of the specified band
// http://localhost:80/bands/bandID/name
func BandName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bandID"]

	n, err := data.GetBandName(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(n, w)
}

// BandMembers returns a list of members in the specified band
// http://localhost:80/bands/bandID/members
func BandMembers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b := vars["bandID"]

	members, err := data.GetBandMembers(b)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(members, w)
}

// BandMember returns information about the specified band member
// http://localhost:80/bands/bandID/members/memberID
func BandMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b := vars["bandID"]
	m := vars["memberID"]

	member, err := data.GetBandMember(b, m)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(member, w)
}

// BandMemberName returns the name of the specified band member
// http://localhost:80/bands/bandID/members/memberID/name
func BandMemberName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b := vars["bandID"]
	m := vars["memberID"]

	name, err := data.GetBandMemberName(b, m)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(name, w)
}

// BandMemberInstruments returns the instruments played by the specified band member
// http://localhost:80/bands/bandID/members/memberID/instruments
func BandMemberInstruments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b := vars["bandID"]
	m := vars["memberID"]

	instruments, err := data.GetBandMemberInstruments(b, m)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(instruments, w)
}

// BandMemberInstrument returns the specified instrument played by the specified band member
// http://localhost:80/bands/bandID/members/memberID/instruments/instruments
func BandMemberInstrument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b := vars["bandID"]
	m := vars["memberID"]
	i := vars["instrumentID"]

	instruments, err := data.GetBandMemberInstrument(b, m, i)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(instruments, w)
}

// BandMemberFounder returns if the specified band member was a founder
// http://localhost:80/bands/bandID/members/memberID/founder
func BandMemberFounder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b := vars["bandID"]
	m := vars["memberID"]

	name, err := data.GetBandMemberFounder(b, m)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(name, w)
}

// BandMemberCurrent returns if the specified band member was a founder
// http://localhost:80/bands/bandID/members/memberID/current
func BandMemberCurrent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b := vars["bandID"]
	m := vars["memberID"]

	name, err := data.GetBandMemberCurrent(b, m)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(name, w)
}

// BandYear returns the year the specified band formed
// http://localhost:80/bands/bandID/year
func BandYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bandID"]

	n, err := data.GetBandYear(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialize(n, w)
}

// Helper functions

// serialize does the JSON marshalling and set the successful HTTP status code
func serialize(v interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// dumpHeader dumps out a response header for debugging
func dumpHeader(h http.Header) {
	for k, v := range h {
		fmt.Printf("[%s]=%s", k, v)
	}
}
