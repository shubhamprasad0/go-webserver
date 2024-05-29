package server

import (
	"bytes"
	"net"
	"testing"
)

func TestNewResponse(t *testing.T) {
	resp := NewResponse(StatusOK)
	if resp.StatusCode != StatusOK {
		t.Errorf("Expected status code %d, got %d", StatusOK, resp.StatusCode)
	}
	if resp.StatusName != "OK" {
		t.Errorf("Expected status name 'OK', got %s", resp.StatusName)
	}
}

func TestSetHeader(t *testing.T) {
	resp := NewResponse(StatusOK)
	resp.SetHeader("Content-Type", "application/json")
	if resp.Headers["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got %s", resp.Headers["Content-Type"])
	}
}

func TestSetData(t *testing.T) {
	resp := NewResponse(StatusOK)
	resp.SetData(`{"key": "value"}`)
	if resp.Data != `{"key": "value"}` {
		t.Errorf("Expected data '%s', got %s", `{"key": "value"}`, resp.Data)
	}
}

func TestMarshal(t *testing.T) {
	resp := NewResponse(StatusOK)
	resp.SetHeader("Content-Type", "application/json")
	resp.SetData(`{"key": "value"}`)
	expectedOutput := "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n{\"key\": \"value\"}"

	output := string(resp.Marshal())
	if output != expectedOutput {
		t.Errorf("Expected marshaled response '%s', got '%s'", expectedOutput, output)
	}
}

func TestSend(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to set up listener: %v", err)
	}
	defer listener.Close()

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Errorf("Failed to accept connection: %v", err)
			return
		}
		defer conn.Close()

		resp := NewResponse(StatusOK)
		resp.SetHeader("Content-Type", "application/json")
		resp.SetData(`{"key": "value"}`)
		resp.Send(conn)
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("Failed to dial listener: %v", err)
	}
	defer conn.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(conn)
	expectedOutput := "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n{\"key\": \"value\"}"
	if buf.String() != expectedOutput {
		t.Errorf("Expected sent response '%s', got '%s'", expectedOutput, buf.String())
	}
}
