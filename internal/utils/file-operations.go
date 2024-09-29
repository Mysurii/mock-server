package utils

import (
	"fmt"
	"os"
	"text/template"
)

type TemplateInfo struct {
	FilePath     string
	TemplatePath string
}

// Generate file from a template
func GenerateFileFromTemplate(ti TemplateInfo) error {
	tmpl, err := template.ParseFiles(ti.TemplatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %v", ti.TemplatePath, err)
	}


	file, err := os.Create(ti.FilePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", ti.FilePath, err)

	}
	defer file.Close()

	err = tmpl.Execute(file, nil)
	if err != nil {
		return fmt.Errorf("failed to execute template %s: %v", ti.TemplatePath, err)
	}

	return nil
}

// Validate that template file can be parsed successfully
func ValidateTemplate(templatePath string) error {
	_, err := template.ParseFiles(templatePath)
	return err
}