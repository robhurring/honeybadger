package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/robhurring/honeybadger"
)

func main() {
	token := os.Getenv("TOKEN")

	honeybadger := honeybadger.New(token)
	results, err := honeybadger.Projects()
	if err != nil {
		panic(err)
	}

	fmt.Printf("You have %d projects in honeybadger!\n", len(results.Results))

	debug(results)
}

func debug(obj ...interface{}) {
	data, _ := json.Marshal(obj)
	fmt.Println(string(data))
}
