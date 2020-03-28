package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var (
	uploadFileKey = "upload-key"
)

func main() {
	url := "http://127.0.0.1:8080/upload"
	path := "/tmp/test.txt"
	params := map[string]string{
		"key1": "val1",
	}
	req, err := NewFileUploadRequest(url, path, params)
	if err != nil {
		fmt.Printf("error to new upload file request:%s\n", err.Error())
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error to request to the server:%s\n", err.Error())
		return
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(body)
}

// NewFileUploadRequest ...
func NewFileUploadRequest(url, path string, params map[string]string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	// > body
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(uploadFileKey, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	for k, v := range params {
		if err := writer.WriteField(k, v); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req, err
}
