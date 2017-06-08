package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/koding/multiconfig"
)

type (
	Server struct {
		Name     string
		Port     int `default:"6066"`
		Enabled  bool
		Users    []string
		Postgres Postgres
	}

	// Postgres is here for embedded struct feature
	Postgres struct {
		Enabled           bool
		Port              int
		Hosts             []string
		DBName            string
		AvailabilityRatio float64
	}
)

func main() {
	m := multiconfig.NewWithPath("conf.toml")

	serverConfig := new(Server)

	m.MustLoad(serverConfig)

	fmt.Println("After Loading: ")
	fmt.Printf("%s+v\n", serverConfig)

	if serverConfig.Enabled {
		fmt.Println("Enabled field is set to true")
	} else {
		fmt.Println("Enabled field is set to false")
	}

	spew.Dump(serverConfig)
}
