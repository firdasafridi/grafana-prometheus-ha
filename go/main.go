package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.String("p", "8080", "help message for flag n")

func main() {

	flag.Parse()
	log.Println("running port: ", *port)
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":"+*port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", *port)
}
