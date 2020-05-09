package main

import (
	"fmt"
	"os"
	"os/exec"
)

const vadDir = "vad"

func vadRun(wavlist string) (string, error) {
	if err := os.Chdir(vadDir); err != nil {
		fmt.Printfln("vad: change dir error")
		return "", err
	}

	cmd := exec.Command("./output/modelvad", "./conf/modelvad.conf", wavlist, "vad.lst")
	fmt.Println(cmd.Path)
	fmt.Println(cmd.Args)
	fmt.Println(cmd.Env)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("vad error: %s", err.Error())
		return "", err
	}

	fmt.Println("vad Done", err)
	return "vad.lst", nil

	file := "/tmp/8315.wav"
	items := filepath.Split(file)
	fmt.Println(items)
	
}

func main() {
	vadDir := "vad"
	os.Chdir(vadDir)
	pwd, err := os.Getwd()
	fmt.Printf("pwd: %s %s\n", pwd, err)
	if pwd != vadDir {
		fmt.Println("change dir error")
	}
	res, err := cmdRun("./wav.list")
	fmt.Println(res)

	file, err := os.Open(res)
}
