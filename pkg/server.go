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
		res := NewResponse(StatusInternalServerError)
		res.Send(conn)
	}
	res := NewResponse(StatusOK)
	res.SetData(fmt.Sprintf("Requested Path: %s", req.Path))
	res.Send(conn)
}
