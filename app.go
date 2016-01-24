package main

type AppContext struct {
}

var context AppContext

func main() {
	e := NewEndpoint("localhost:8080", "example.com.crt", "example.com.key")

	e.OriginServers = []*originServer{
		&originServer{
			Host:  "requestb.in",
			HTTPS: false,
		},
	}
	e.ProxyStats = NewProxyStats(10, 1)
	e.ProxyStats.Track()

	go e.Serve()

	e2 := NewEndpoint("localhost:8081", "example.com.crt", "example.com.key")

	os, _ := NewOriginServer("http2.golang.org", 80, false)
	e2.OriginServers = []*originServer{os}

	go e2.Serve()

	select {}
}
