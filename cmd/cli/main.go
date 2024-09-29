package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/mysurii/mock-server/internal/server"
	"github.com/mysurii/mock-server/internal/utils"
)

func main() {
	filePath := flag.String("file", "mock/server-config.json", "Path to the JSON file")
	exampleFlag := flag.Bool("example", false, "Create example mock files")
	flag.Parse() // Parse the command-line flags

	if *exampleFlag {
		baseDir := "mock"
		generateExampleFolders(baseDir)
		generateExampleFiles(baseDir)
	}

	s := server.New(*filePath)
	s.StartServer()

}

func generateExampleFolders(baseDir string) {

	dirs := []string{
		filepath.Join(baseDir, "auth"),
		filepath.Join(baseDir, "users"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}
}

func generateExampleFiles(baseDir string) {
	templates := prepareTemplates(baseDir)

	for _, tmplInfo := range templates {
		err := utils.GenerateFileFromTemplate(tmplInfo)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func prepareTemplates(baseDir string) []utils.TemplateInfo {
	return []utils.TemplateInfo{
		{FilePath: filepath.Join(baseDir, "server-config.json"), TemplatePath: "internal/templates/config.json.templ"},
		{FilePath: filepath.Join(baseDir, "auth", "login.json"), TemplatePath: "internal/templates/login.json.templ"},
		{FilePath: filepath.Join(baseDir, "users", "users.json"), TemplatePath: "internal/templates/users.json.templ"},
		{FilePath: filepath.Join(baseDir, "users", "user.json"), TemplatePath: "internal/templates/user.json.templ"},
	}
}
