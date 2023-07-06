package main

import "github.com/bachiitter/api/app"

func main() {
	err := app.New()
	if err != nil {
		panic(err)
	}
}
