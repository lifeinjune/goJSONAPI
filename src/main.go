package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", router))

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("still alive")
}
