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
	log.Println("Server started")
}

func echo(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	response := buf.String()
	log.Println("Received:", response)
	io.WriteString(w, response)
}
