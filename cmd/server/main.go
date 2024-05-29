package main

import (
	"flag"

	s "github.com/shubhamprasad0/go-webserver/pkg"
)

func main() {
	configFile := flag.String("config", "config/conf.yaml", "server config file")
	flag.Parse()

	config := s.FromYaml(*configFile)
	server := s.New(config)
	server.Start()
}
