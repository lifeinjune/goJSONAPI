package main

import (
	"encoding/json"          //encoding and decoding of json
	"fmt"                    //to use Println
	"github.com/gorilla/mux" //router
	"log"                    //log server exit
	"net/http"               //http package
	"github.com/gorilla/handlers"
)

func main() {
	router := mux.NewRouter() //make new multiplexer
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/message", handleQryMessage).Methods("GET") //match pattern with handler
	router.HandleFunc("/m/{msg}", handleURLMessage).Methods("GET") // take path parameters (/m/{msg})
	// same as http.HandleFunc(), pattern and handler function passed
	// as argument. if request match the path, handler function called
	// and passing (http.ResponseWriter, *http.Request) as argument to
	//handler function.
	// matches request with HTTP get method
	// HandleFunc receives handler func(w http.ResponseWriter, req *http.request)
	// in this case helthCheck meet the requirment as argument

	headerOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET","POST","OPTIONS"})

	fmt.Println("Running server!") //write onto the terminal
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router))
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

func handleQryMessage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()      //returns value of URL query (url? after values)
	message := vars.Get("msg") //returns first value of the key this case "msg"

	json.NewEncoder(w).Encode(map[string]string{"message": message})
	//respond client with json format of map
}

func handleURLMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)    //returns pattern(route) variable this case {msg:(request variable after/m/)}
	message := vars["msg"] //put value of msg to variable message

	json.NewEncoder(w).Encode(map[string]string{"message": message})
	//respons back to client json format of map which is object
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	//http.ResponseWriter used to respond to the client
	//http.Request give the information about request
	json.NewEncoder(w).Encode("still alive")
	//json.NewEncoder returns encoder pointer writes to w
	// this case writer is http.ResponseWriter
	// then Encoder() json encoding with argument passed to it
	// this case "still alive"
}
