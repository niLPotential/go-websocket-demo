package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
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

	cw, err := c.Writer(ctx, websocket.MessageText)
	if err != nil {
		log.Printf("Error creating writer: %v", err)
		return
	} else {
		log.Println("Writer created")
	}

	cw.Write([]byte(fmt.Sprintf(`<div id="time">%d</div>`, time.Now().Unix())))
}
