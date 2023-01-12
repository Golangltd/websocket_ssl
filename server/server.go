package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

type WSServer struct {
	ListenAddr string
}

// 处理
func (this *WSServer) handler(conn *websocket.Conn) {
	fmt.Printf("a new ws conn: %s->%s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	var err error
	for {
		var reply string
		err = websocket.Message.Receive(conn, &reply)
		if err != nil {
			fmt.Println("receive err:", err.Error())
			break
		}
		fmt.Println("Received from client: " + reply)
		if err = websocket.Message.Send(conn, reply); err != nil {

			fmt.Println("send err:", err.Error())
			break
		}
	}
}
func (this *WSServer) start() error {
	http.Handle("/ws", websocket.Handler(this.handler))
	fmt.Println("begin to listen")
	//err := http.ListenAndServe(this.ListenAddr, nil)
	err := http.ListenAndServeTLS(this.ListenAddr, "certificate/ssl.crt", "certificate/ssl.key", nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
		return err
	}
	fmt.Println("start end")
	return nil
}
func main() {
	addr := flag.String("a", "127.0.0.1:12345", "websocket server listen address")
	flag.Parse()
	wsServer := &WSServer{
		ListenAddr: *addr,
	}
	wsServer.start()
	fmt.Println("------end-------")
}
