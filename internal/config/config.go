package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mysurii/mock-server/internal/models"
)

func LoadApiFile(jsonPath string) (models.API, error) {
	raw, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var api models.API

	err = json.Unmarshal(raw, &api)
	if err != nil {
		log.Fatal(" ", err)
	}

	checkDuplicates(api)

	return api, nil
}

func checkDuplicates(api models.API) {
    seen := make(map[string]int)

    for _, endpoint := range api.Endpoints {
        key := fmt.Sprintf("%s:%s", endpoint.Method, endpoint.Path)
        seen[key]++

        // Check if it appeared more than once
        if seen[key] > 1 {
            log.Fatalf("Duplicate endpoint found: Method: %s, Path: %s\n", endpoint.Method, endpoint.Path)
        }
    }
}