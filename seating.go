package main

import (
	"fmt"
	"net/http"
	"code.google.com/p/go.net/websocket"
)

func main() {
	chart := seatingChart("seating-chart")
	arrivals := getInitialArrivals("arrivals")
	ServeAndProtect(":12345", chart, arrivals)
}

func ServeAndProtect(tcpPort string, chart Chart, arrivals []Guest) {
	man := NewArrivalsManager(chart, arrivals)
	man.Run()

	sockHandler := SocketHandler{man}
	http.Handle("/html", http.HandlerFunc(sendHtml))
	http.Handle("/client.js", http.HandlerFunc(sendJavascript))
	http.Handle("/live-feed", websocket.Handler(sockHandler.handle))
	err := http.ListenAndServe(tcpPort, nil)
	if err != nil {
		fmt.Println("\nListenAndServe: " + err.Error())
	}
	fmt.Println("\nThat's all, folks.")
}

func sendHtml(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, ClientHTML)
}

func sendJavascript(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	fmt.Fprint(w, ClientJS)
}

type SocketHandler struct{*ArrivalsManager}

func (handler SocketHandler) handle(ws *websocket.Conn) {
	registrar := handler.Hire(NewRegistrar())
	registrar.Use(NewBrowser(ws))
}

