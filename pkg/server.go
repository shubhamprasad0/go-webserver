package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
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
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	req, err := ExtractRequestData(conn)
	if err != nil {
		res := NewResponse(StatusInternalServerError)
		res.Send(conn)
		return
	}

	path := req.Path
	if path == "/" {
		path = "/index.html"
	}

	requestedFile, err := s.validatePath(path)
	if err != nil {
		if errors.Is(err, ErrUnauthorized) {
			res := NewResponse(StatusUnauthorized)
			res.Send(conn)
			return
		} else {
			res := NewResponse(StatusInternalServerError)
			res.Send(conn)
			return
		}
	}

	data, err := os.ReadFile(requestedFile)
	if err != nil {
		if os.IsNotExist(err) {
			res := NewResponse(StatusNotFound)
			res.Send(conn)
		} else {
			res := NewResponse(StatusInternalServerError)
			res.Send(conn)
		}
		return
	}
	res := NewResponse(StatusOK)
	res.SetHeader("Content-Type", "text/html")
	res.SetData(string(data))
	res.Send(conn)
}

func (s *Server) validatePath(path string) (string, error) {
	requestedFile, err := filepath.Abs(fmt.Sprintf("%s%s", s.Config.RootPath, path))
	if err != nil {
		return "", err
	}
	rootPath, err := filepath.Abs(s.Config.RootPath)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(requestedFile, rootPath) {
		return "", ErrUnauthorized
	}
	return requestedFile, nil
}
