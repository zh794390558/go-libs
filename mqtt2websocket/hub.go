package main

import (
	log "github.com/cihub/seelog"
	"github.com/zh794390558/go-libs/mqtt"
	"strings"
)

type HubMsg struct {
	topic string
	msg   []byte
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients. topic -> map[*Client]bool
	clients map[string]map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan HubMsg

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	mqttClients map[string]*mqtt.Client
}

func newHub() *Hub {
	hub := &Hub{
		//broadcast:  make(chan []byte),
		broadcast:  make(chan HubMsg),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]map[*Client]bool),
        mqttClients: make(map[string]*mqtt.Client),
	}
	go hub.run()
	return hub
}

//发送文本到指定的房间的client
func sendMsgToClient(msg []byte, topic string, hub *Hub) {
	for client, ok := range hub.clients[topic] {
		if !ok {
			continue
		}
		select {
		case client.send <- msg:
			log.Infof("client_send:%s, client:%s", topic, (*client).conn.RemoteAddr())
		default:
			log.Infof("send failed, unregister:%s", (*client).conn.RemoteAddr())
			close(client.send)
			delete(hub.clients[topic], client)
		}
	}
}

func newMqttClient(clientId string) *mqtt.Client {
	clientId = strings.Join([]string{*name, clientId}, "-")
	client := mqtt.NewClient(
		clientId,
		strings.Join([]string{*mqttAddr, *mqttPort}, ":"),
		*mqttUser,
		*mqttPasswd)
	err := client.Connect()
	if err != nil {
		log.Error(err)
		return nil
	}
	return client
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			if _, ok := h.clients[client.topic]; ok {
				stockClients := h.clients[client.topic]
				stockClients[client] = true
				h.clients[client.topic] = stockClients
			} else {
				h.clients[client.topic] = map[*Client]bool{
					client: true,
				}
			}
			log.Infof("register_topic:%s, room_size:%d, register:%s", client.topic, len(h.clients[client.topic]), (*client).conn.RemoteAddr())

			if _, ok := h.mqttClients[client.topic]; !ok {
				mqttCli := newMqttClient(*name)
				h.mqttClients[client.topic] = mqttCli
				if err := mqttCli.Subscribe(
					func(c *mqtt.Client, msg string) {
						log.Infof("mqtt msg: %v", msg)
						h.broadcast <- HubMsg{client.topic, []byte(msg)}
					}, 2, client.topic); err != nil {
					log.Infof("ws %s : mqtt subscribe error: %v", client.conn.RemoteAddr(), err)
				}
			}

		case client := <-h.unregister:
			log.Infof("logout_topic:%s, before room_size:%d, unregister_:%s", client.topic, len(h.clients[client.topic]), (*client).conn.RemoteAddr())

			if _, ok := h.clients[client.topic]; ok {
				stockClients := h.clients[client.topic]
				delete(stockClients, client)
				close(client.send)
				h.clients[client.topic] = stockClients

				// all clients  are closed, then delte map info
				if len(stockClients) == 0 {
					delete(h.clients, client.topic)
					h.mqttClients[client.topic].Unsubscribe(client.topic)
					delete(h.mqttClients, client.topic)
				}
			} else {
				close(client.send)
			}

			//how to delete mqtt client

		case message := <-h.broadcast:
			sendMsgToClient(message.msg, message.topic, h)
			log.Infof("broad cast msg: %s to topic %s", message.msg, message.topic)
		}
	}
}
