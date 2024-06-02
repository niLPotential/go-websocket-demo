package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var Tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", DemoHandler)
	http.HandleFunc("/ws", WebSocketHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetTime() []byte {
	return []byte(fmt.Sprintf(`<div id="time">%d</div>`, time.Now().Unix()))
}
