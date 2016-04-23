package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":8080", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	log.Println("Server started")
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	response := buf.String()
	log.Println("Received:", response)
	io.WriteString(w, response)
}
