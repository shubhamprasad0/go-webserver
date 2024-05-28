package main

import s "github.com/shubhamprasad0/go-webserver/pkg"

func main() {
	config := s.DefaultConfig()
	server := s.New(config)
	server.Start()
}
