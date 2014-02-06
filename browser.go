package main

import "code.google.com/p/go.net/websocket"
import "fmt"


type Browser struct {
	sendOn chan Guest
	listenOn chan interface{}
}

func FromBrowser(b *Browser) chan Guest {
	return b.sendOn
}

func ToBrowser(b *Browser) chan interface{} {
	return b.listenOn
}

func NewBrowser(ws *websocket.Conn) *Browser{
	var ready string
	websocket.Message.Receive(ws, &ready)
	if ready !=  "ok" {
		fmt.Println("The ping was", ready)
		return nil
	}
	
	var b = new(Browser)
	b.sendOn = make(chan Guest)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err == nil {
				FromBrowser(b) <- unpack(msg)
				continue
			}
			break
		}
		close(FromBrowser(b))
	}()

	b.listenOn = make(chan interface{})
	go func() {
		for obj := range ToBrowser(b) {
			websocket.JSON.Send(ws, obj)
		}
		ws.Close()
	}()
	return b
}


