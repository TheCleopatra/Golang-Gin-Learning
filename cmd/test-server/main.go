package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Service is listening on port 8000. In golang ;)")
	log.Fatal(http.ListenAndServe(":8000", nil))
}