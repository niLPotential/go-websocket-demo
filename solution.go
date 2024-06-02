package main

import (
	"context"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func SolHandler(w http.ResponseWriter, r *http.Request) {
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

		// create new writer every loop
		cw, err := c.Writer(ctx, websocket.MessageText)
		if err != nil {
			log.Printf("Error creating writer: %v", err)
			return
		} else {
			log.Println("Writer created")
		}
		// defer cw.Close() // defer won't work

		if _, err = cw.Write(GetTime()); err != nil {
			log.Printf("Error writing message: %v", err)
		} else {
			log.Println("Message successfully written")
		}
		cw.Close() // close manually after write
	}
}
