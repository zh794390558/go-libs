package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	const jsonStream = `
	[
	{"Name": "Ed", "Text": "Knock knock."},
	{"Name": "Sam", "Text": "Who's there?"},
	{"Name": "Ed", "Text": "Go fmt."},
	{"Name": "Sam", "Text": "Go fmt who?"},
	{"Name": "Ed", "Text": "Go fmt yourself!"}

	]
	`
	type Message struct {
		Name, Text string
	}

	var m []Message
	if err := json.Unmarshal([]byte(jsonStream), &m); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
	fmt.Println("vim-go")
}
