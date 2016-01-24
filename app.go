package main

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

	e.Serve()
}
