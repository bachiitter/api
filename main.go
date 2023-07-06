package main

import (
	"github.com/bachiitter/api/app"
	"github.com/bachiitter/api/utils"
)

func main() {

	// load env
	err := utils.LoadENV()
	if err != nil {
		panic(err)
	}

	// init server
	app.New()

}
