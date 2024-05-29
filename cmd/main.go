package main

import s "github.com/shubhamprasad0/go-webserver/pkg"

func main() {
	config := s.DefaultConfig()
	config.RootPath = "test/www"
	server := s.New(config)
	server.Start()
}
