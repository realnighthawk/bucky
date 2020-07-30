package web

type Request struct {
	Meta string      `json:"meta,omitempty"`
	Body interface{} `json:"body,omitempty"`
}

type Response struct {
	Code string      `json:"code,omitempty"`
	Body interface{} `json:"body,omitempty"`
}
