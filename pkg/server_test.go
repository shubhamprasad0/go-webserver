package server

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestValidatePath(t *testing.T) {
	config := DefaultConfig()
	config.RootPath = "./test_data"
	server := New(config)

	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"/index.html", "test/index.html", nil},
		{"/../secrets.txt", "", ErrUnauthorized},
	}

	for _, test := range tests {
		result, err := server.validatePath(test.input)
		if test.err != nil && !errors.Is(err, test.err) {
			t.Errorf("Expected error %v, got %v", test.err, err)
		} else if result != "" && strings.HasSuffix(result, test.expected) {
			t.Errorf("Expected path %s, got %s", test.expected, result)
		}
	}
}

func TestHandleConnection(t *testing.T) {
	config := DefaultConfig()
	config.RootPath = "./test_data"
	server := New(config)

	go server.Start()
	time.Sleep(10 * time.Millisecond)

	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", server.Config.Port))
	if err != nil {
		t.Fatalf("Failed to dial listener: %v", err)
	}
	defer conn.Close()

	request := "GET /index.html HTTP/1.1\r\nHost: localhost\r\n\r\n"
	conn.Write([]byte(request))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	expectedResponse := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n<html><body>Hello, World!</body></html>"
	if string(buf[:n]) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(buf[:n]))
	}
}

func setupTestFiles() {
	os.MkdirAll("./test_data", 0755)
	os.WriteFile("./test_data/index.html", []byte("<html><body>Hello, World!</body></html>"), 0644)
}

func teardownTestFiles() {
	os.RemoveAll("./test_data")
}

func TestMain(m *testing.M) {
	setupTestFiles()
	code := m.Run()
	teardownTestFiles()
	os.Exit(code)
}
