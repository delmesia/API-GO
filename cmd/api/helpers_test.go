package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func healthcheckHandlerMarshalIndent(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": "development",
		"version":     "1.0.0",
	}

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func healthcheckHandlerMarshal(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": "development",
		"version":     "1.0.0",
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func BenchmarkMarshalIndent(b *testing.B) {
	w := httptest.NewRecorder()
	r := new(http.Request)
	for n := 0; n < b.N; n++ {
		healthcheckHandlerMarshalIndent(w, r)
	}
}

func BenchmarkMarshal(b *testing.B) {
	w := httptest.NewRecorder()
	r := new(http.Request)
	for n := 0; n < b.N; n++ {
		healthcheckHandlerMarshal(w, r)
	}
}
