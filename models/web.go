package models

type Request struct {
	Meta string      `json:"meta,omitempty"`
	Body interface{} `json:"body,omitempty"`
}

type Response struct {
	Code string      `json:"code,omitempty"`
	Body interface{} `json:"body,omitempty"`
}

type Health struct {
	Version string `json:"version"`
	Status  string `json"status"`
	Error   string `json:"error"`
}
