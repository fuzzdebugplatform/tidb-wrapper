package main

import (
	"example/protocol"
	"net/http"
)

func main() {
	http.HandleFunc("/test", protocol.Handle())
	http.ListenAndServe(":4333", nil)
}