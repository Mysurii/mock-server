package models

type Endpoint struct {
	Method   string  `json:"method"`
	Status   int     `json:"status"`
	Path     string  `json:"path"`
	JsonPath *string `json:"jsonPath,omitempty"`
}

type API struct {
	Host      string     `json:"host"`
	Port      int        `json:"port"`
	Endpoints []Endpoint `json:"endpoints"`
}