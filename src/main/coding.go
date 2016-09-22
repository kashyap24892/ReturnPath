package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Country ... Set of All Countries
type Country []string

//States ... Map of string to string array
//States... Maps Country to corresponding cities
type States map[string][]string

//Cities ... Map of string to string array
//Cities... Maps State to corresponding cities
type Cities map[string][]string

var country Country
var states States
var cities Cities

//HTMLHandler ... reads HTML file from filesystem
func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string
		path := r.URL.Path[1:]
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "text/javascript"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("File Not Found "))
	}
}

// JSHandler1 ... handles ajax requests
func JSHandler1(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		r.ParseForm()
		fmt.Fprintf(w, strings.Join(states[string(r.Form.Get("Country"))], ","))
	}
}

// JSHandler2 ... handles ajax requests
func JSHandler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		r.ParseForm()
		fmt.Fprintf(w, strings.Join(cities[string(r.Form.Get("State"))], ","))
	}
}

func main() {
	initialize()
	http.HandleFunc("/", HTMLHandler)
	http.HandleFunc("/states", JSHandler1)
	http.HandleFunc("/cities", JSHandler2)
	http.ListenAndServe("localhost:1337", nil)
}

func initialize() {
	country = []string{"United States", "India"}

	states = make(map[string][]string)
	states["United States"] = append(states["United States"], "Texas")
	states["United States"] = append(states["United States"], "New York")

	cities = make(map[string][]string)
	cities["Texas"] = append(cities["Texas"], "Dallas")
	cities["Texas"] = append(cities["Texas"], "Austin")
	cities["Texas"] = append(cities["Texas"], "Houston")

	cities["New York"] = append(cities["New York"], "NYC")
	cities["New York"] = append(cities["New York"], "Rochester")
	cities["New York"] = append(cities["New York"], "Syracuse")

}
