package server

import (
	"log"
	"net/http"
)

// ResponseRecorder is a custom ResponseWriter to capture the response status.
type ResponseRecorder struct {
	http.ResponseWriter
	Status int
}

func (lrw *ResponseRecorder) WriteHeader(code int) {
	lrw.Status = code
	lrw.ResponseWriter.WriteHeader(code)
}


// LoggingMiddleware logs incoming requests and their responses.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a ResponseRecorder to capture the response
		recorder := &ResponseRecorder{ResponseWriter: w}
		
		// Call the next handler
		next.ServeHTTP(recorder, r)

		// Log the response status
		log.Printf("%v %v - %v", r.Method, r.URL.Path, recorder.Status)
	})
}

