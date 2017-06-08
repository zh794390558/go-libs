package main

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
)

type Config struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time
}

func main() {

	var conf Config

	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(conf)

}
