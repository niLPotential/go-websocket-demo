package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", DemoHandler)
	http.HandleFunc("/ws", WebSocketHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
