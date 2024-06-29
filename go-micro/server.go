package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloHandlerResponse struct {
	Message    string `json:"message"`
	Occupation string `json:"occupation"`
	ID         int    `json:"id"`
}

func main() {
	port := 2323

	http.HandleFunc("/helloworld", helloHandler)

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	res := helloHandlerResponse{
		Message:    "Willkommen",
		Occupation: "Cloud Engineer",
		ID:         203511,
	}
	data, err := json.Marshal(res)
	if err != nil {
		panic("Told you so!")
	}
	fmt.Fprint(w, string(data))
}
