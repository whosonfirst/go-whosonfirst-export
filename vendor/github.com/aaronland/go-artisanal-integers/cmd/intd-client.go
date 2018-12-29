package main

import (
	"github.com/aaronland/go-artisanal-integers/application"
	"log"
	"os"
)

func main() {

	flags := application.NewClientApplicationFlags()

	app, err := application.NewClientApplication()

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(flags)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
