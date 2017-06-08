package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	//main operation modes
	write = flag.Bool("w", false, "write result back instaed of stdout\n\t\tDefault: NOwrite back")

	// layout conrol

	tabWidth = flag.Int("tabwidth", 8, "tab width\n\t\tDefault: Standard")

	//debugging
	cpuprofile = flag.String("cputprofile", "", "write cpu profile to thie file\n\t\tDefault: no default")
)

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage: %s [flags] file [path...]\n\n",
		"CommandLineFlag")

	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	fmt.Printf("Before parsing the flags\n")
	fmt.Printf("T: %d\nW: %s\nC: '%s'\n",
		*tabWidth, strconv.FormatBool(*write), *cpuprofile)

	flag.Usage = usage
	flag.Parse()

	// there is also a mandatory non-flag arguments
	if len(flag.Args()) < 1 {
		usage()
	}

	fmt.Printf("Testing the flag pcakage\n")
	fmt.Printf("T: %d\nW: %s\nC: '%s'\n",
		*tabWidth, strconv.FormatBool(*write), *cpuprofile)

	for index, element := range flag.Args() {
		fmt.Printf("I: %d C: '%s'\n",
			index, element)
	}
}
