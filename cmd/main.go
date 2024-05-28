package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicf("Failed to start server: %+v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to handle connection: %+v", err)
			continue
		}
		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	// Read data from connection
	var data []byte
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF || err == net.ErrClosed {
				break
			}
		}
		data = append(data, buf[:n]...)
		if n < 1024 {
			break
		}
	}
	lines := strings.Split(string(data), "\n")
	firstLine := lines[0]
	path := strings.Split(firstLine, " ")[1]
	conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\nRequested path: %s\r\n", path)))
}
