package main

import (
	"context"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func FailHandler(w http.ResponseWriter, r *http.Request) {
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

	cw, err := c.Writer(ctx, websocket.MessageText)
	if err != nil {
		log.Printf("Error creating writer: %v", err)
		return
	} else {
		log.Println("Writer created")
	}
	defer cw.Close()

	if _, err = cw.Write(TimeDiv); err != nil {
		log.Printf("Error writing message: %v", err)
	}
}