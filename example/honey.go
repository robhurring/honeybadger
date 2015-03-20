package main

import (
	"fmt"
	"os"

	"github.com/robhurring/honeybadger"
)

func main() {
	// get your token on the https://app.honeybadger.io/users/edit page
	token := os.Getenv("HONEYBADGER_API_TOKEN")
	honeybadger := honeybadger.New(token)

	projects, err := honeybadger.Projects()
	if err != nil {
		panic(err)
	}

	fmt.Printf("You have %d projects in honeybadger!\n", len(projects.Results))
}
