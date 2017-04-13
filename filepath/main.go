package main

import (
	"fmt"
	"os"

	"path/filepath"
)

func main() {
	fmt.Println(filepath.Base("./main.go"))
	fmt.Println(filepath.Ext("./main.go"))
	fmt.Println(filepath.Abs("./main.go"))
	fmt.Println(os.Hostname())
}
