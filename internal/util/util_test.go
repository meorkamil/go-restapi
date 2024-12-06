package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test ConfigInit
func TestUtilConfigInit(t *testing.T) {
	file := "../../config/config.yml"
	config := ConfigInit(file)
	assert.NotNil(t, config)
}

// Test Decode
func TestUtilDecode(t *testing.T) {
	req := &http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(`{"name": "kamil", "age": 31}`))),
	}

	var result interface{}
	err := json.NewDecoder(req.Body).Decode(&result)

	if err != nil {
		t.Error(err)
	}
}

// Test Encode
type TestResponseWriter struct {
	HeaderMap  map[string][]string
	StatusCode int
	Body       bytes.Buffer
}

func (m *TestResponseWriter) Header() http.Header {
	return m.HeaderMap
}

func (m *TestResponseWriter) Write(b []byte) (int, error) {
	return m.Body.Write(b)
}

func (m *TestResponseWriter) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}

func TestUtilEncode(t *testing.T) {
	// Create sample http response
	Headermap := make(map[string][]string)
	Headermap = map[string][]string{
		"Content-Type":  {"application/json"},
		"Accept":        {"application/json", "text/html"},
		"Authorization": {"Bearer token123"},
	}

	testResponse := TestResponseWriter{
		HeaderMap: Headermap,
	}

	err := Encode(&testResponse, nil, http.StatusOK, "Success")
	if err != nil {
		t.Error(err)
	}
}
