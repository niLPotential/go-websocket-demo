package main

import (
	"context"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("Error accepting WebSocket handshake: %v", err)
		return
	} else {
		log.Printf("WebSocket connected: %v", r.URL)
	}
	defer c.CloseNow()

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	for {
		if _, _, err = c.Read(ctx); err != nil {
			log.Printf("Error reading message: %v", err)
		} else {
			log.Println("Message successfully read")
		}

		if err = c.Write(ctx, websocket.MessageText, GetTime()); err != nil {
			log.Printf("Error writing message: %v", err)
		} else {
			log.Println("Message successfully written")
		}
	}
}
