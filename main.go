package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleHome(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]interface{}{
		"status":  200,
		"message": "OK",
	})
}
