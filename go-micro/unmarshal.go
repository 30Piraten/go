package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type handlerRes struct {
	Text    string `json:"text"`
	Address string `json:"address"`
}

type handlerReq struct {
	User string `json:"user"`
	City string `json:"city"`
}

func main() {
	port := 2324

	http.HandleFunc("/", reqHandler)
	log.Printf("Serving at 127.0.0.1:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		http.Error(w, "Empty body", http.StatusBadRequest)
		return
	}

	var request handlerReq
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	res := handlerRes{
		Text:    "Wie gehts " + request.User,
		Address: "Wo wohnst du? " + request.City,
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}

// server must be running first
// go run unmarshal.go

// using curl to connect via terminal
//curl -X POST -H "Content-Type: application/json" -d '{"user":"spence99", "city":"Berlin"}' http://127.0.0.1:2324/
