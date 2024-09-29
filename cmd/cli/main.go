package main

import (
	"flag"

	"github.com/mysurii/mock-server/internal/server"
)

const (
	charsetUTF8 = "charset=UTF-8"
)

func main() {
	// Step 2: Use flag to get the file path from the user
	filePath := flag.String("file", "db.json", "Path to the JSON file")
	flag.Parse() // Parse the command-line flags

	s := server.New(*filePath)
	s.StartServer(8080)

}