package server

import (
	"net"
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	tests := []struct {
		line           string
		expectedMethod string
		expectedPath   string
		expectError    bool
	}{
		{"GET / HTTP/1.1", "GET", "/", false},
		{"POST /submit HTTP/1.1", "POST", "/submit", false},
		{"INVALID", "", "", true},
	}

	for _, test := range tests {
		method, path, err := parseRequestLine(test.line)
		if test.expectError && err == nil {
			t.Errorf("Expected error for line %q, got nil", test.line)
		} else if !test.expectError {
			if err != nil {
				t.Errorf("Unexpected error for line %q: %v", test.line, err)
			} else if method != test.expectedMethod || path != test.expectedPath {
				t.Errorf("Expected method %q and path %q, got method %q and path %q", test.expectedMethod, test.expectedPath, method, path)
			}
		}
	}
}

func TestParseHeaderLine(t *testing.T) {
	tests := []struct {
		line          string
		expectedKey   string
		expectedValue string
		expectError   bool
	}{
		{"Content-Type: text/html", "Content-Type", "text/html", false},
		{"X-Custom-Header: custom_value", "X-Custom-Header", "custom_value", false},
		{"InvalidHeader", "", "", true},
	}

	for _, test := range tests {
		key, value, err := parseHeaderLine(test.line)
		if test.expectError && err == nil {
			t.Errorf("Expected error for line %q, got nil", test.line)
		} else if !test.expectError {
			if err != nil {
				t.Errorf("Unexpected error for line %q: %v", test.line, err)
			} else if key != test.expectedKey || value != test.expectedValue {
				t.Errorf("Expected key %q and value %q, got key %q and value %q", test.expectedKey, test.expectedValue, key, value)
			}
		}
	}
}

func TestExtractRequestData(t *testing.T) {
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
		request := "GET /test HTTP/1.1\r\nHost: localhost\r\n\r\n"
		conn.Write([]byte(request))
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("Failed to dial listener: %v", err)
	}
	defer conn.Close()

	req, err := ExtractRequestData(conn)
	if err != nil {
		t.Fatalf("Failed to extract request data: %v", err)
	}

	if req.Method != "GET" {
		t.Errorf("Expected method GET, got %s", req.Method)
	}

	if req.Path != "/test" {
		t.Errorf("Expected path /test, got %s", req.Path)
	}

	expectedHeaders := map[string]string{
		"Host": "localhost",
	}

	for k, v := range expectedHeaders {
		if req.Headers[k] != v {
			t.Errorf("Expected header %s: %s, got %s", k, v, req.Headers[k])
		}
	}
}
