package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {

	filepath := "../one-8k.pcm"
	uploadFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer uploadFile.Close()

	var buff bytes.Buffer
	writer := multipart.NewWriter(&buff)
	writer.WriteField("field", "this is a field")
	w, _ := writer.CreateFormFile("file", filepath)
	io.Copy(w, uploadFile)

	writer.Close()

	var client http.Client
	resp, err := client.Post("http://127.0.0.1:8080/upload", writer.FormDataContentType(), &buff)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("client:", string(data))
}
