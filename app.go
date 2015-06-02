package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("Starting error: ", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello CF!")
	fmt.Fprintln(w, "host: ", os.Getenv("VCAP_APP_HOST"))
	fmt.Fprintln(w, "port: ", os.Getenv("VCAP_APP_PORT"))
}
