package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world!\n")
}
func ByeServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Bye, Bye!\n")
}
func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/bye", ByeServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
