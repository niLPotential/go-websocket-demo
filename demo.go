package main

import (
	"context"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	Tmpl.Execute(w, nil)
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("Error accepting WebSocket handshake: %v", err)
		return
	} else {
		log.Println("WebSocket connected")
	}
	defer c.CloseNow()

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	// c.Write(ctx, websocket.MessageText, TimeDiv)

	cw, err := c.Writer(ctx, websocket.MessageText)
	if err != nil {
		log.Printf("Error creating writer: %v", err)
		return
	} else {
		log.Println("Writer created")
	}
	defer cw.Close()

	cw.Write(TimeDiv)
}
