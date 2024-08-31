package response

import (
	"encoding/json"
	"net/http"

	"log/slog"
)

type Reciever struct {
	http.ResponseWriter

	headers    http.Header
	statusCode int
	body       []byte
}

func NewReciever() *Reciever {
	return &Reciever{
		headers:    make(http.Header),
		statusCode: 200,
		body:       []byte{},
	}
}

func (d *Reciever) Header() http.Header {
	return d.headers
}

func (d *Reciever) WriteHeader(statusCode int) {
	d.statusCode = statusCode
}

func (d *Reciever) Write(body []byte) (int, error) {
	d.body = append(d.body, body...)
	return len(body), nil
}

func (d *Reciever) ToJSON() string {
	response := map[string]interface{}{
		"statusCode": d.statusCode,
		"headers":    d.headers,
		"body":       string(d.body),
	}

	jsonReciever, err := json.Marshal(response)
	if err != nil {
		slog.Error("failed to marshal response to JSON", "error", err)
		return "{}"
	}

	return string(jsonReciever)
}
