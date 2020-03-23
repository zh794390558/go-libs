package mqtt

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMqtt(t *testing.T) {
	var (
		clientId = "pibigstar"
		Host     = "127.0.0.1:1883"
		UserName = "pibigstar"
		Password = "123456"
		wg       sync.WaitGroup
	)
	client := NewClient(clientId, Host, UserName, Password)
	err := client.Connect()
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println("new client")

	wg.Add(1)
	go func() {
		err := client.Subscribe(func(c *Client, msg *Message) {
			fmt.Printf("recv msg: %+v \n", msg)
			wg.Done()
		}, 1, "mqtt")

		if err != nil {
			panic(err)
		}
	}()
	fmt.Println("new pub")

	msg := &Message{
		ClientID: clientId,
		Type:     "text",
		Data:     "Hello Pibistar",
		Time:     time.Now().Unix(),
	}
	data, _ := json.Marshal(msg)

	err = client.Publish("mqtt", 1, false, data)
	if err != nil {
		panic(err)
	}
	fmt.Println("sub msg")

	wg.Wait()
}
