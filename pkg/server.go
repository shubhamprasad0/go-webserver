package server

import (
	"fmt"
	"log"
	"net"
)

const (
	BufferSize uint16 = 1024
)

type Server struct {
	Config *Config
}

func New(config *Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.Port))
	if err != nil {
		log.Panicf("Failed to start server: %+v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to handle connection: %+v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	req, err := ExtractRequestData(conn)
	if err != nil {
		conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
	}
	conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\nRequested path: %s\r\n", req.Path)))
}
