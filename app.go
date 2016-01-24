package main

import "log"

type AppContext struct {
	Stats proxyStats
}

var context AppContext

func main() {
	e := &endpoint{}

	e.Address = "localhost:8080"
	e.Certs = certPair{"example.com.crt", "example.com.key"}

	e.OriginServers = []*originServer{
		&originServer{
			Host:  "requestb.in",
			HTTPS: false,
		},
	}
	e.ProxyStats = NewProxyStats(10, 1)
	e.ProxyStats.Track()

	go e.Serve()

	e2 := &endpoint{}
	e2.Address = "localhost:8081"
	e2.Certs = certPair{"example.com.crt", "example.com.key"}

	os, _ := NewOriginServer("http2.golang.org", 80, false)
	log.Println("os:", os)
	e2.OriginServers = []*originServer{os}

	go e2.Serve()

	select {}
}
