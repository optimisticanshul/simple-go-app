package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Response is just a very basic example.
type Response struct {
	Status string `json:"status,omitempty"`
}

// GetStatus returns always the same response.
func GetStatus(w http.ResponseWriter, _ *http.Request) {
	b := Response{Status: "Hello world!"}
	json.NewEncoder(w).Encode(b)
}

// GetDetectify returns always the same response.
func GetDetectify(w http.ResponseWriter, _ *http.Request) {
	b := Response{Status: "Hello World again"}
	json.NewEncoder(w).Encode(b)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetStatus).Methods("GET")
	router.HandleFunc("/detectify", GetDetectify).Methods("GET")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
