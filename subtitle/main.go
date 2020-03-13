package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"

	log "github.com/cihub/seelog"
	mqtt "github.com/zh794390558/go-libs/mqtt"
)

type AsrContentStruct struct {
	Nbest     []string `json:"nbest"`
	Uncertain []string `json:"uncertain"`
}

type AsrParamStruct struct {
	Cpid   int    `json:"cpid"`
	ErrNo  int    `json:"err_no"`
	Idx    int    `json:"idx"`
	Sid    string `json:"sid"`
	Status int    `json:"status"`
}

type AsrInfoStruct struct {
	AsrDeviceid string `json:"asr_deviceid"`
}

type InDTO struct {
	AsrContent  AsrContentStruct `json:"asr_content"`
	AsrDeviceid string           `json:"asr_deviceid"`
	AsrParam    AsrParamStruct   `json:"asr_param"`
	RoomID      string           `json:"room_id"`
}

type DataStruct struct {
	AsrContent AsrContentStruct `json:"asr_content"`
	AsrParam   AsrParamStruct   `json:"asr_param"`
	AsrInfo    AsrInfoStruct    `json:"asr_info"`
}

type OutDTO struct {
	RoomID string     `json:"room_id"`
	Data   DataStruct `json:"data"`
	ErrMsg string     `json:"err_msg"`
	ErrNo  int        `json:"err_no"`
}

var addr = flag.String("addr", ":8082", "http service address")

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
	log.Infof("realtime_subtitle start on port%s", *addr)
	flag.Parse()

	clientId := "subtitle"
	client := mqtt.NewClient(clientId, "127.0.0.1:1883", "test", "12345")
	err = client.Connect()
	if err != nil {
		log.Error(err)
		return
	}

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

	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("token") != "54321" {
			log.Error("wrong token")
			w.Write([]byte("error"))
			return
		}
		s, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error(err)
			w.Write([]byte("error"))
			return
		}
		log.Info(string(s))

		var inDTO InDTO
		err = json.Unmarshal(s, &inDTO)
		if err != nil {
			log.Error(err)
			w.Write([]byte("error"))
			return
		}

		outDTO := OutDTO{
			RoomID: inDTO.RoomID,
			Data: DataStruct{
				AsrParam:   inDTO.AsrParam,
				AsrContent: inDTO.AsrContent,
				AsrInfo: AsrInfoStruct{
					AsrDeviceid: inDTO.AsrDeviceid,
				},
			},
			ErrNo: 0,
		}
		data, err := json.Marshal(outDTO)
		if err != nil {
			log.Error(err)
			w.Write([]byte("error"))
			return
		}

		err = client.Publish("mqtt", 1, false, data)
		if err != nil {
			log.Error(err)
			w.Write([]byte("error"))
			return
		}

		w.Write([]byte("ok"))
		return
	})

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Critical("ListenAndServe: ", err)
	}
}
