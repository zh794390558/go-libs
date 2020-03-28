package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var (
	// 文件 key
	uploadFileKey = "upload-key"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error to start http server:%s", err.Error())
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 接受文件
	file, header, err := r.FormFile(uploadFileKey)
	if err != nil {
		// ignore the error handler
	}
	log.Printf("selected file name is %s", header.Filename)
	// 将文件拷贝到指定路径下，或者其他文件操作
	dst, err := os.Create(header.Filename)
	if err != nil {
		// ignore
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		// ignore
	}
}
