package main

import (
	"log"

	"main.go/internal/app"
	"main.go/internal/config"
	"main.go/internal/flags"
)

func main() {
	flags, err := flags.ParseOptions()
	if err != nil {
		log.Fatal(err)
	}

	config, err := config.FromFile(*flags)
	if err != nil {
		log.Fatal(err)
	}

	cataas := app.New(*config)

	if err = cataas.GetCat(); err != nil {
		log.Fatal(err)
	}

	if err = cataas.SavePicture(); err != nil {
		log.Fatal(err)
	}
}
