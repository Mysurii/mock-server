package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mysurii/mock-server/internal/models"
	"github.com/mysurii/mock-server/internal/utils"
)

// List of allowed HTTP methods
var allowedMethods = []string{
    "GET",
    "POST",
    "PUT",
	"PATCH",
    "DELETE",
}

func (s *server) registerRoutes() *http.ServeMux {
	println("registering routes:")

	mux := http.NewServeMux()

	for _, endpoint := range s.api.Endpoints {
		if !isAllowedMethod(endpoint.Method) {
			log.Printf("Method %s not allowed - skipping endpoint", endpoint.Method)
			continue
		}

		mux.HandleFunc(fmt.Sprintf("%s /api%s", endpoint.Method, endpoint.Path), handler(&endpoint))
		log.Printf("%s: %s\n", endpoint.Method, endpoint.Path)
	}

	return mux
}

func handler (endpoint *models.Endpoint) http.HandlerFunc {	
	return func(w http.ResponseWriter, r *http.Request) {
		println("JSONPATH")
		println(endpoint.JsonPath)
		if endpoint.JsonPath == nil {
			WriteSuccess(w, endpoint.Status, nil)
			return
		}

		println("SECOND")

		payload, err := utils.LoadPayload(*endpoint.JsonPath)
		if err != nil {
			WriteError(w, http.StatusInternalServerError, fmt.Sprintf("Could not open file %s. Are you sure it exists?", *endpoint.JsonPath))
			return
		}
		
		WriteSuccess(w, endpoint.Status, &payload)
	}
}

// Function to check if a method is allowed
func isAllowedMethod(method string) bool {
    for _, allowed := range allowedMethods {
        if allowed == method {
            return true // Method is allowed
        }
    }
    return false // Method is not allowed
}


// WriteSuccess writes a success message as JSON to the http.ResponseWriter.
func WriteSuccess(w http.ResponseWriter, status int, data interface{}) {
	response := map[string]interface{}{
		"statusCode": status,
		"success": true,
	}

	// Only add the data field if data is not nil
	if data != nil {
		response["data"] = data
	}

	WriteJSON(w, status, response)
}

// WriteError writes an error message as JSON to the http.ResponseWriter.
func WriteError(w http.ResponseWriter, status int, message string) {
	response := map[string]interface{}{
		"error":   message,
		"statusCode": status,
		"success": false,
	}
	WriteJSON(w, status, response)
}


// WriteJSON writes the given data as JSON to the http.ResponseWriter
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
    // Set the content type to application/json
    w.Header().Set("Content-Type", "application/json")
    
    // Write the HTTP status code
    w.WriteHeader(status)

    // Encode the data into JSON and write it to the response
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
    }
}