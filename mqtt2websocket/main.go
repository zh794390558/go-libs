package main

import (
	"flag"
	log "github.com/cihub/seelog"
	"net/http"
	"strings"
)

var name = flag.String("name", "wsbroadcast-live", "app name")
var addr = flag.String("addr", ":8085", "http service address")

var mqttAddr = flag.String("mqtt-host", "127.0.0.1", "mqtt host")
var mqttPort = flag.String("mqtt-port", "1883", "mqtt port")
var mqttUser = flag.String("mqtt-user", "subtitle-live", "mqtt username")
var mqttPasswd = flag.String("mqtt-passwd", "1234", "mqtt password")
var topicPrefix = flag.String("topic-prefix", "/subtitle/live", "mqtt topic prefix")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Info(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	logger, err := log.LoggerFromConfigAsFile("conf/seelog.xml")
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)
	defer logger.Flush()

	flag.Parse()
	log.Infof("wsbroadcast start on port%s", *addr)

	hub := newHub()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, "home.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		room_id := r.URL.Query().Get("room_id")
		lang := r.URL.Query().Get("lang") // en, zh ...
		log.Info("url: ", r.URL)
		topic := strings.Join([]string{*topicPrefix, lang, room_id}, "/")
		log.Info("watch topic: ", topic)
		serveWs(hub, topic, w, r)
	})

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Critical("ListenAndServe: ", err)
	}
}
