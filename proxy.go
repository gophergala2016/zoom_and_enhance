package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"code.google.com/p/go-uuid/uuid"
)

// Generate a unique request ID so that incoming requests can be associated with
// responding requests to origin server.
func genID() string {
	id := uuid.NewRandom()
	log.Println("UUID:", id.String())

	return id.String()
}

func (e *endpoint) wildcardHandler(w http.ResponseWriter, r *http.Request) {
	e.ProxyStats.Insert(1)

	//log.Printf("Req: %#v", r)
	c := http.Client{}

	req, err := translateRequest(r, e)
	if err != nil {
		log.Println("Request Error:", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Println("Error:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Response Error:", err)
	}
	w.Write(body)
}

func (e *endpoint) translateURL(r *http.Request) string {
	newURL := url.URL{}
	newURL.Host = e.OriginServers[0].Host
	newURL.Path = r.URL.Path
	newURL.Scheme = "http"
	if e.OriginServers[0].HTTPS {
		newURL.Scheme = "https"
	}

	return newURL.String()
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func setHandlers(r *http.ServeMux, e *endpoint) {
	r.HandleFunc("/", e.wildcardHandler)
}

func translateRequest(r *http.Request, e *endpoint) (*http.Request, error) {

	url := e.translateURL(r)

	req, err := http.NewRequest(r.Method, url, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
