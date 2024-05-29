package server

import (
	"bytes"
	"fmt"
	"net"
)

type StatusCode uint

const (
	StatusOK                  StatusCode = 200
	StatusInternalServerError StatusCode = 500
	StatusNotFound            StatusCode = 404
	StatusUnauthorized        StatusCode = 401
)

var StatusToName = map[StatusCode]string{
	StatusOK:                  "OK",
	StatusInternalServerError: "Internal Server Error",
	StatusNotFound:            "Not Found",
	StatusUnauthorized:        "Unauthorized",
}

type Response struct {
	StatusCode StatusCode
	StatusName string
	Headers    map[string]string
	Data       string
}

func NewResponse(statusCode StatusCode) *Response {
	statusName := StatusToName[statusCode]

	return &Response{
		StatusCode: statusCode,
		StatusName: statusName,
		Headers:    make(map[string]string),
	}
}

func (r *Response) SetHeader(key, value string) {
	r.Headers[key] = value
}

func (r *Response) SetData(data string) {
	r.Data = data
}

func (r *Response) Marshal() []byte {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.StatusCode, r.StatusName))
	for k, v := range r.Headers {
		b.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	b.WriteString("\r\n")
	if r.Data != "" {
		b.WriteString(r.Data)
	}
	return b.Bytes()
}

func (r *Response) Send(conn net.Conn) {
	data := r.Marshal()
	conn.Write(data)
}
