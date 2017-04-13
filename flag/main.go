package main

import (
	"flag"
	"fmt"
	"github.com/topicai/candy"
)

func main() {
	var flagvar int
	// define flag
	var ip = flag.Int("flagname", 1234, "help message for flagname")
	flag.IntVar(&flagvar, `adfb`, 23, "hdkfjsdlfj")

	//parse the command line
	flag.Parse()

	fmt.Println("ip has value", *ip)
	fmt.Println("flagvar has value", flagvar)
	fmt.Println(flag.NArg(), flag.Arg(0))

	fmt.Println(candy.GoPath())
}
