package main

import (
	"encoding/json"          //to create json response
	"fmt"                    //to use Println
	"github.com/gorilla/mux" //router
	"log"                    //log server exit
	"net/http"               //http package
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
