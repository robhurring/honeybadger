package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/robhurring/honeybadger"
)

func main() {
	token := os.Getenv("TOKEN")
	// projectId, _ := strconv.Atoi(os.Getenv("ID"))

	honeybadger := honeybadger.New(token)
	results, err := honeybadger.Projects()
	if err != nil {
		panic(err)
	}

	// debug(results)
	fmt.Printf("You have %d projects in honeybadger!\n", len(results.Results))
}

func debug(obj ...interface{}) {
	data, _ := json.Marshal(obj)
	fmt.Println(string(data))
}
