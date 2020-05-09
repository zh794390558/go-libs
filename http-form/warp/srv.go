package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"bytes"
	"os/exec"
)

var task chan string = make(chan string)
var taskDone chan string = make(chan string)

func upload(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(w, "unexpected error: %s\n", err)
			break
		}

		filename := part.FileName()
		if filename == "" {
			fmt.Println("part is not a file: key = ", part.FormName())
			continue
		}
		filename = filepath.Base(filename)

		data, err := ioutil.ReadAll(part)
		if err != nil {
			fmt.Fprintln(w, err)
			continue
		}
		fmt.Printf("srv: FormName(%s),FileName(%s)=>%s\n", part.FormName(), part.FileName())

		// 检查后缀
		ext := strings.ToLower(filepath.Ext(filename))
		if ext != ".pcm" {
			w.Write([]byte(fmt.Sprintf("only support pcm: %s", ext)))
			return
		}

		// 保存
		os.Mkdir("./uploaded", 0777)
		savePath := filepath.Join("./uploaded", filename)
		fmt.Println("save to :", savePath)
		saveFile, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		saveFile.Write(data)
		defer saveFile.Close()

		// vad
		task <- savePath
		
		<- taskDone

		w.Write([]byte("OK"))
	}
}

func main() {
	go func() {
		 pcmPath := <- task
		 fmt.Println("vad proc: ", pcmPath)

		 buf := bytes.NewBuffer(nil)
		 cmd := exec.Command("bash", "-c", "ls", "-lth", ".")
		 cmd.Stdout = buf
		 err := cmd.Run()
		 if err != nil{
			 fmt.Printf("run cmd [%s %s] : err %s ", cmd.Path, cmd.Args, err)
		 }
		 fmt.Printf("vad: %s %s : \n %s \n", cmd.Path, cmd.Args, buf.String())
		 fmt.Println("vad Done")

		 taskDone <- "Done"
	}()

	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}