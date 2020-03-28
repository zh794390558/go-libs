package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type HttpClient struct {
	client  http.Client
	url     string
	pid     string
	appname string
	rate    string
	idx     int

	sid string
}

type AsrParam struct {
	ErrNo  int    `json:"err_no"`
	Err    string `json:"error"`
	Idx    int    `json:"idx"`
	Sid    string
	Status string
}

type Resp struct {
	Param AsrParam `json:"asr_param"`
}

func newHttpClient(url string) *HttpClient {
	return &HttpClient{
		client:  http.Client{Timeout: 10 * time.Second},
		url:     url,
		pid:     "50010",
		appname: "com.test.8k",
		rate:    "8",
		idx:     0,
	}
}

func (h *HttpClient) Url() url.Values {
	v := url.Values{}
	v.Add("pid", h.pid)
	v.Add("appname", h.appname)
	v.Add("rate", h.rate)
	//log.Println("url.Values", v.Encode())
	//s := fmt.Sprintf("%s?%s", h.url, v.Encode())
	//log.Println("url: ", s)
	return v
}

type Info struct {
	Pfm    string
	Devid  string
	Ver    string
	Gender string `json:"gender_recog"`
}

func (h *HttpClient) CreateSession() string {
	jsonStr, _ := json.Marshal(Info{Pfm: "mac", Devid: "xx", Ver: "1", Gender: "1"})

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("create-session", string(jsonStr))
	// !!! very important, not using `defer writer.Close()` , or lossing data , e.g. 1 byte
	writer.Close()

	v := h.Url()
	v.Add("idx", strconv.Itoa(0))
	url := fmt.Sprintf("%s?%s", h.url, v.Encode())
	log.Println("url: ", url)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}

	req.Header.Add("content-type", writer.FormDataContentType())
	req.Header.Add("content-type", "text/json")
	log.Println(req)
	defer req.Body.Close()

	resp, err := h.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(result))

	var r Resp
	json.Unmarshal(result, &r)
	if r.Param.ErrNo != 0 {
		log.Fatalf("create session err: %d, %s", r.Param.ErrNo, r.Param.Err)
	}
	h.sid = r.Param.Sid
	return r.Param.Sid
}

func (h *HttpClient) Get() (response string) {
	v := h.Url()
	v.Add("sid", h.sid)
	url := fmt.Sprintf("%s?%s", h.url, v.Encode())
	log.Println("Get url: ", url)

	rsp, err := h.client.Get(url)
	defer rsp.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	result, _ := ioutil.ReadAll(rsp.Body)

	response = string(result)
	return response
}

func (h *HttpClient) Post(idx int, data []byte, filename string) (content string) {
	v := h.Url()
	v.Add("sid", h.sid)
	v.Add("idx", strconv.Itoa(idx))
	url := fmt.Sprintf("%s?%s", h.url, v.Encode())
	log.Println("Post url: ", url)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	//part, err := writer.CreateFormFile("asr_audio", filename)
	part, err := writer.CreateFormField("asr_audio")
	if err != nil {
		return ""
	}
	//_ = writer.WriteField("create-session", string(jsonStr))
	part.Write(data)
	// !!! very important, not using `defer writer.Close()` , or lossing data , e.g. 1 byte
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("content-type", writer.FormDataContentType())
	req.Header.Add("content-type", "audio/pcm")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	//log.Println(req)

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

func (h *HttpClient) PostFile(filename string) (content string) {
	v := h.Url()
	v.Add("sid", h.sid)
	v.Add("idx", strconv.Itoa(-1))
	url := fmt.Sprintf("%s?%s", h.url, v.Encode())
	log.Println("Post file url: ", url)

	openFile, err := os.Open(filename)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer openFile.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("asr_audio", filename)
	//part, err := writer.CreateFormField("asr_audio")
	if err != nil {
		return ""
	}
	//_ = writer.WriteField("create-session", string(jsonStr))
	//part.Write(data)
	io.Copy(part, openFile)
	// !!! very important, not using `defer writer.Close()` , or lossing data , e.g. 1 byte
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("content-type", writer.FormDataContentType())
	req.Header.Add("content-type", "audio/pcm")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	log.Println(req)

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

func main() {
	h := newHttpClient("http://speech.com/asr")
	sid := h.CreateSession()
	log.Println("sessid: ", sid)

	go func() {
		for {
			rsp := h.Get()
			log.Println("get rsp:", rsp)
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("post file")
	datapath := "test-8k.pcm"
	h.PostFile(datapath)

	fmt.Println("post file by chunck")

	file, err := os.Open(datapath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()

	idx := 0
	const bufsize int = 3200
	var buffer []byte = make([]byte, bufsize)
	for {
		n, err := file.Read(buffer[0:])
		log.Printf("Read %d bytes", n)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err.Error())
			break
		}

		idx += 1
		if n < bufsize {
			idx *= -1
			for i := 0; i < (bufsize - n); i++ {
				buffer[n+i] = 0
			}
		}
		log.Printf("Write %d bytes", len(buffer))
		rsp := h.Post(idx, buffer, datapath)
		log.Println("post rsp:", rsp)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		_, _, _ = reader.ReadLine()
	}
}
