package http

import (
	"encoding/json"
	"testing"
	"log"
)

type Msg struct {
	Text string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type NmtData struct {
	Src string `json:"source"`
	Tgt string `json:"target"`
	Text string `json:"translation"`
}

type Result struct {
	Data NmtData `json:"data"`
	Code int `json:"code"`
	TraceId string `json:"trace_id"`
}

func TestHttpClient(t *testing.T){
	client := newHttpClient("zh",  "https://fanyi/api")

	data := Msg{
		Text: "Hello world",
		Source: "en",
		Target: "zh",
	}

	datastr, err := json.Marshal(data)
	if err != nil {
		t.Error("json error")
	}
	log.Println("input", string(datastr))

	rsp := client.Post(data, "application/json")
	log.Println("test rsult" , rsp)

	res := new(Result)
	json.Unmarshal([]byte(rsp), res)
	log.Println(res.Code)
	log.Println(res.Data.Text)
	log.Println(res.TraceId)
}
