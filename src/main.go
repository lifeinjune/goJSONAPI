package main

import (
	"encoding/json"          //encoding and decoding of json
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
	// HandleFunc receives handler func(w http.ResponseWriter, req *http.request)
	// in this case helthCheck meet the requirment as argument

	fmt.Println("Running server!") //write onto the terminal
	log.Fatal(http.ListenAndServe(":3000", router))
	//http.ListenAndServe(addr string, handler Handler) error
	/*'
	listens on TCP network address "addr" and call Serve with
	handler to handle request connection
	accepted connection are configured to enable TCP keep-alives
	handler is typically nil, with that case, DefaultServeMux is used
	always return non-nil error
	*/
	// Fatal same as Print() follow by os.Exit(1)

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("still alive")
	//json.NewEncoder returns encoder pointer writes to w
	// this case writer is http.ResponseWriter
	// then Encoder() json encoding with argument passed to it
	// this case "still alive"
}
