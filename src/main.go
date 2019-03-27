package main

import (
	"encoding/json"          //to create json response
	"fmt"                    //to use Println
	"github.com/gorilla/mux" //router
	"log"                    //log server exit
	"net/http"               //http package
)

func main() {
	router := mux.NewRouter() //make new multiplexer
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	// same as http.HandleFunc(), pattern and handler function passed
	// as argument. if request match the path, handler function called
	// and passing (http.ResponseWriter, *http.Request) as argument to
	//handler function.
	// matches request with HTTP get method

	fmt.Println("Running server!") //write onto the terminal
	log.Fatal(http.ListenAndServe(":3000", router))

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("still alive")
}
