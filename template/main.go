package main

import (
	"fmt"
	"os"
	"text/template"
)

type podMeta struct {
	PodName       string
	ContainerName string
	Image         string
	Cmds          []string
	Args          []string
	Cpu           string
	Mem           string
	Gpu           uint8
	NfsServer     string
	NfsSrcPath    string
	NfsDstPath    string
}

func main() {
	m := podMeta{
		PodName:       "test",
		ContainerName: "test",
		Image:         "bootstrapper:5000/zhanghui/uers:v1.0",
		Cmds: []string{
			"/bin/bash",
			"-c",
		},
		Args: []string{
			"sleep 10",
		},
		Cpu:        "450m",
		Mem:        "1Gi",
		Gpu:        2,
		NfsServer:  "10.10.10.251",
		NfsSrcPath: "/volume1/gfs",
		NfsDstPath: "/gfs",
	}

	//work
	t := template.Must(template.ParseFiles("pod.templ"))
	err := t.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	//work
	t = template.Must(template.New("pod.templ").ParseFiles("pod.templ"))
	err = t.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	//work
	t = template.Must(template.New("test").ParseFiles("pod.templ"))

	fmt.Println(t.Name(), len(t.Templates()))

	err = t.ExecuteTemplate(os.Stdout, "pod.templ", m)
	if err != nil {
		panic(err)
	}

	for tmp := range t.Templates() {
		fmt.Println(tmp)
	}

	fmt.Println(t.Name(), len(t.Templates()))

	fmt.Println()

	// can not work
	t = template.Must(template.New("test").ParseFiles("pod.templ"))
	err = t.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}

	fmt.Println("name", t.Name())

}
