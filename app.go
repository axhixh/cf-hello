package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting web server at %s", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Starting error: ", err)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Headers:")
	for k, v := range req.Header {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello CF!")
	fmt.Fprintln(w, "See /headers for request headers")
}
