package main

import (
	"encoding/json"
	"net/http"
)

type post struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Body         string `json:"-"` // Won't Marshal it
	LikeCount    int
	CommentCount int
	IsVerified   bool
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", rootHandler)
	handler.HandleFunc("/json", jsonHandler)
	handler.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", handler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	p := post{
		ID:           "sdjfw9rjnbfigu48t934jgh934t",
		Title:        "My Post",
		Body:         "The posts's description",
		LikeCount:    1233,
		CommentCount: 25,
		IsVerified:   true,
	}

	bytes, _ := json.Marshal(p)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
