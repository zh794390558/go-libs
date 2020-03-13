package http 

import (
	"net/http"
	"io"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"
	"log"
)


type HttpClient struct {
	client http.Client
	langType string
	url string
}

func newHttpClient(langType string, url string) *HttpClient {
	return &HttpClient{
		client: http.Client{Timeout: 10 * time.Second},
		langType : langType,
		url : url,
	}
}

func (h *HttpClient) Get() (response string) {
	rsp, err := h.client.Get(h.url)
	defer rsp.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return "" 
	}

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := rsp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Println(err.Error())
			return "" 
		}
	}

	response = result.String()
	return response
}

func (h *HttpClient) Post(data interface{}, contentType string) (content string) {
	jsonStr, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", h.url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	req.Header.Add("apikey", "")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	resp, err := h.client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}
