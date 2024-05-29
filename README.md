# Go Web Server

## Overview

This is a simple HTTP server implemented in Go implementing a very small subset of HTTP/1.1 just for learning purposes. It serves static files from a specified root directory and handles basic HTTP requests. This project includes the following features:

- Parsing HTTP requests
- Constructing and sending HTTP responses
- Serving static files
- Configurable server settings through a YAML file

## Getting Started

### Prerequisites

- Go 1.22 or later

### Installation

1. Clone the repository:

    ```sh
    git clone git@github.com:shubhamprasad0/go-webserver.git
    cd go-webserver
    ```

2. Build the application using the Makefile:

    ```sh
    make build
    ```

    This will compile the application and place the binary in the `bin` directory.

### Running the Server

After building the application, you can run the server with the following command:

```sh
./bin/server -config config/conf.yaml
```

By default, the server will start on port `8080` and serve files from the `test/www` directory.

### Configuration

The server configuration is handled through a YAML file. By default, the configuration file is located at `config/conf.yaml`. You can specify a different configuration file using the `-config` flag.

Example `config/conf.yaml`:
```yaml
port: 8080
root_path: "test/www"
```

### Testing

To run the unit tests, use the following command:

```sh
go test ./pkg/...
```

This will execute all the tests in the `pkg` directory.

## Makefile Commands

- `make build`: Compiles the application and places the binary in the `bin` directory.
- `make clean`: Removes the `bin` directory and the compiled binary.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

