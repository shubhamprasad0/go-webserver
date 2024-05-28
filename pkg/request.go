package server

import (
	"io"
	"net"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    string
}

func readDataFromRequest(conn net.Conn) []byte {
	var data []byte
	buf := make([]byte, BufferSize)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF || err == net.ErrClosed {
				break
			}
		}
		data = append(data, buf[:n]...)
		if n < int(BufferSize) {
			break
		}
	}
	return data
}

func ExtractRequestData(conn net.Conn) (*Request, error) {
	data := readDataFromRequest(conn)
	segments := strings.Split(string(data), "\r\n")
	var requestLine string
	headers := make(map[string]string)
	bodyStart := -1
	for i, line := range segments {
		if i == 0 {
			requestLine = line
			continue
		}
		if line == "" {
			bodyStart = i + 1
			break
		}
		key, val, err := parseHeaderLine(line)
		if err != nil {
			return nil, err
		}
		headers[key] = val
	}
	body := strings.Join(segments[bodyStart:], "\r\n")

	method, path, err := parseRequestLine(requestLine)
	if err != nil {
		return nil, err
	}

	return &Request{
		Method:  method,
		Path:    path,
		Headers: headers,
		Body:    body,
	}, nil

}

func parseRequestLine(rl string) (string, string, error) {
	segments := strings.Split(rl, " ")
	if len(segments) != 3 {
		return "", "", ErrMalformedRequest
	}
	method := segments[0]
	path := segments[1]
	return method, path, nil
}

func parseHeaderLine(hl string) (string, string, error) {
	segments := strings.SplitN(hl, ":", 2)
	if len(segments) != 2 {
		return "", "", ErrMalformedRequest
	}
	key := strings.TrimSpace(segments[0])
	val := strings.TrimSpace(segments[1])
	return key, val, nil
}
