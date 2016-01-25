package main

import (
	"errors"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

type endpoint struct {
	OriginServers []*originServer
	Features      features
	Address       string
	Certs         certPair
	ProxyStats    *proxyStats
}

type certPair struct {
	CertFile string
	KeyFile  string
}

type features map[string]bool

func NewEndpoint(addr string, cert string, key string) *endpoint {
	return &endpoint{
		Address: addr,
		Certs:   certPair{cert, key},
	}
}

func (e *endpoint) Serve() error {
	// TODO: Currently only support single server per endpoint
	if len(e.OriginServers) > 1 {
		return errors.New("Currently only 1 origin server per endpoint is supported.")
	}

	// Create http2 server
	var srv http.Server

	// Set Port
	srv.Addr = e.Address

	// Mux setup
	router := http.NewServeMux()

	// Set mux
	srv.Handler = router

	// Set handlers
	setHandlers(router, e)

	err := http2.ConfigureServer(&srv, &http2.Server{})
	if err != nil {
		return err
	}

	log.Printf("Listening on %s", srv.Addr)
	srv.ListenAndServeTLS(e.Certs.CertFile, e.Certs.KeyFile)

	return nil
}
