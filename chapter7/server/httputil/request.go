package httputil

import (
	"encoding/json"
	"net/http"
	"strings"
)

// SerializableRequest implements all the methods of http.Request
// but can serialize to JSON.
type SerialzableRequest struct {
	*http.Request
}

// ToJSON serializes the current object and returns an string.
func (sr *SerialzableRequest) ToJSON() string {
	data, _ := json.Marshal(sr.serialize())
	return string(data)
}

type serializedHeader struct {
	Key   string
	Value string
}

type serializedRequest struct {
	Method  string
	Host    string
	Query   string
	Path    string
	Headers []serializedHeader
}

func (sr *SerialzableRequest) serialize() serializedRequest {
	var headers []serializedHeader
	for k, v := range sr.Header {
		headers = append(headers, serializedHeader{Key: k, Value: strings.Join(v, ",")})
	}

	return serializedRequest{
		Host:    sr.Host,
		Path:    sr.URL.Path,
		Query:   sr.URL.RawQuery,
		Method:  sr.Method,
		Headers: headers,
	}
}
