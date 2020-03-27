package main

import (
	"fmt"
	"os"

	"path/filepath"
)

func main() {
	fmt.Println("Base: ", filepath.Base("/tmp/main.go"))
	fmt.Println("Ext: ", filepath.Ext("/tmp/main.go"))
	fmt.Println(filepath.Abs("/tmp/main.go"))
	fmt.Println(os.Hostname())

	filename := "tmp.go"
	file := filepath.Join("./uploaded/", filename)
	fmt.Println(file)
}
