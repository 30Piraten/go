package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	port := 2323

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port: %v\n", 2323)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Willkomen!\n")
}
