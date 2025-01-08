package animal

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type HttpStatus struct {
	Status string
}

type Message struct {
	Message string
}

func EncodeToJson(w http.ResponseWriter, v any, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func EncodeToXml(w http.ResponseWriter, v any, code int) {
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.WriteHeader(code)
	_, _ = fmt.Fprint(w, xml.Header)
	_ = xml.NewEncoder(w).Encode(v)
}

func Encode(w http.ResponseWriter, r *http.Request, v any, code int) {
	accept := r.Header.Get("Accept")
	if accept == "application/xml" {
		EncodeToXml(w, v, code)
	} else {
		EncodeToJson(w, v, code)
	}
}

func Decode(v any, r *http.Request) error {
	accept := r.Header.Get("Accept")
	if accept == "application/xml" {
		return xml.NewDecoder(r.Body).Decode(v)
	} else {
		return json.NewDecoder(r.Body).Decode(v)
	}
}
