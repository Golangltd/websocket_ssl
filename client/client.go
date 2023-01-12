package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

func timeWriter(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 2)
		websocket.Message.Send(conn, "hello world")
	}
}

func main() {
	wsConfig, _ := websocket.NewConfig("wss://127.0.0.1:12345/ws", "http://localhost/")
	wsConfig.TlsConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "www.example2.com",
	}
	ws, err := websocket.DialConfig(wsConfig)
	if err != nil {
		log.Fatal(err)
	}

	message := []byte("hello, world!")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)
	go timeWriter(ws)
	for {
		var msg = make([]byte, len(message))
		_, err = ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receive: %s\n", msg)
	}
}
