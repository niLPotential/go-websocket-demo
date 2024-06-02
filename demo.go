package main

import "net/http"

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
