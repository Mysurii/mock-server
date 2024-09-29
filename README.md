# Mock Server Generator

A lightweight mock server generator in Go that allows developers to easily create mock APIs based on JSON configurations, facilitating rapid frontend development and testing without backend dependencies.

## Features

- **Dynamic API Endpoints**: Define API endpoints in JSON configuration files.
- **Flexible Response Handling**: Return mock responses based on the defined structure in the configuration.
- **File-Based Configuration**: Load payloads from JSON files for endpoints.
- **CLI Integration**: Simple command-line interface to initialize and run the mock server.

## Installation

To install the mock server generator, clone the repository and build the project:

```bash
git clone https://github.com/Mysurii/mock-server.git
cd mock-server-generator
go build ./cmd/cli/main.go
```

## Usage

To run the mock server, use the following command:

```bash
./main.exe --config path/to/config.json

Or run the following command to create needed config file & example payload files

./main.exe --example
```

### Config format

The configuration file should be in JSON format. Here is an example:

```bash
{
  "port": 3000,
  "endpoints": [
    {
      "method": "POST",
      "status": 200,
      "path": "/login",
      "jsonPath": "./mock/auth/login.json"
    },
    {
      "method": "GET",
      "status": 200,
      "path": "/users",
      "jsonPath": "./mock/users/users.json"
    },
    {
      "method": "POST",
      "status": 201,
      "path": "/users"
    },
    {
      "method": "GET",
      "status": 200,
      "path": "/users/{id}",
      "jsonPath": "./mock/users/user.json"
    }
  ]
}
```

## Contribution

Contributions are welcome! Feel free to submit a pull request or create an issue for any enhancements or bug fixes.
