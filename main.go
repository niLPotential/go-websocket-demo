package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var Tmpl = template.Must(template.ParseFiles("templates/index.html"))
var TimeDiv = []byte(fmt.Sprintf(`<div id="time">%d</div>`, time.Now().Unix()))

func main() {
	http.HandleFunc("/", DemoHandler)
	http.HandleFunc("/ws", WebSocketHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
