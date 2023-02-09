package main

import (
	"log"

	"main.go/interal/app"
	"main.go/interal/config"
	"main.go/interal/flags"
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

	cataas, err := app.New(*config)
	if err != nil {
		log.Fatal(err)
	}

	err = cataas.SavePicture()
	if err != nil {
		log.Fatal(err)
	}

	defer cataas.Res.Body.Close()
}
