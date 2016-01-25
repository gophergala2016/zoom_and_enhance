package main

import "errors"

type originServer struct {
	Host  string
	Port  int
	HTTPS bool
}

func NewOriginServer(hostname string, port int, https bool) (*originServer, error) {
	if hostname == "" {
		return nil, errors.New("Hostname is invalid")
	}

	if port >= 65535 || port < 1 {
		return nil, errors.New("Port is invalid")
	}

	return &originServer{hostname, port, https}, nil
}
