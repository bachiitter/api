package main

import (
	"github.com/bachiitter/api/app"
)

func main() {

	// init server
	err := app.New()
	if err != nil {
		panic(err)
	}

}
