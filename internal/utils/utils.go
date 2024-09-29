package utils

import (
	"encoding/json"
	"fmt"
	"os"
)



func LoadPayload(jsonPath string) (interface{}, error) {
	raw, err := os.ReadFile(jsonPath)

	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	// Initialize a variable to hold the unmarshaled data
	var payload interface{}

	// Unmarshal the JSON data into the payload variable
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	return payload, nil
}