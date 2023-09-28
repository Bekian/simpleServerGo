package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTP!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	log.Println("Server started on localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
