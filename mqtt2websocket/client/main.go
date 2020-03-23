package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

var addr = flag.String("addr", "10.90.66.56:8085", "http service address")
var wg sync.WaitGroup

func NewWebSocketConn(url string, clientID int) {
	defer wg.Done()
	ws, err := websocket.Dial(url, "", "http://127.0.0.1:8080/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[%d] is connected\n", clientID)
	for {
		var msg [512]byte
		_, err := ws.Read(msg[:]) //此处阻塞，等待有数据可读
		if err != nil {
			fmt.Printf("[%d] read error: %s\n", clientID, err)
			return
		}
		// if clientID == 1 {
		// 	fmt.Printf("[%d] received: %s\n", clientID, msg)
		// }
	}
}

func main() {
	flag.Parse()
	fmt.Println("start")
	url := "ws://" + *addr + "/ws"
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 10)
		go NewWebSocketConn(url, i)
	}
	wg.Wait()
	fmt.Println("exit test")
}
