package main

// Response ...
type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
	Error   error       `json:"error"`
}

// NewResponse ...
func NewResponse(status int, payaload interface{}, e error) *Response {
	r := &Response{}

	if status != nil {
		r.Status = status
	}

	if payload != nil {
		r.Payload = payaload
	}

	if e != nil {
		r.Error = e
	}

	return r
}
