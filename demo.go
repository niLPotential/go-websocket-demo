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

	if err = c.Write(ctx, websocket.MessageText, GetTime()); err != nil {
		log.Printf("Error writing message: %v", err)
	}

	// cw, err := c.Writer(ctx, websocket.MessageText)
	// if err != nil {
	// 	log.Printf("Error creating writer: %v", err)
	// 	return
	// } else {
	// 	log.Println("Writer created")
	// }
	// defer cw.Close()

	// cw.Write(GetTime())
}
