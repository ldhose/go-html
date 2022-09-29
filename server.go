package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})

	fmt.Printf("Starting server at port 8085")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
