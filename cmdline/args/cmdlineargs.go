package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("i am", os.Args[0])

	baseName := filepath.Base(os.Args[0])

	fmt.Println("The base name is", baseName)

	fmt.Println("Argument # is ", len(os.Args))

	if len(os.Args) > 1 {
		fmt.Println("the first command line argument is ", os.Args[1])
	}

}
