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
	http.HandleFunc("/demo", TmplHandler)
	http.HandleFunc("/demo/ws", DemoHandler)

	http.HandleFunc("/fail", TmplHandler)
	http.HandleFunc("/fail/ws", FailHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TmplHandler(w http.ResponseWriter, r *http.Request) {
	Tmpl.Execute(w, r.URL)
}
