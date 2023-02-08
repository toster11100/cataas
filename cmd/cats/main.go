package main

import (
	"io"
	"log"
	"os"

	"main.go/interal/app"
	"main.go/interal/config"
	"main.go/interal/flags"
)

func main() {
	flags, err := flags.New()
	if err != nil {
		log.Fatal(err)
	}

	defVal, err := config.DefConfig(flags.Config)
	if err != nil {
		log.Fatal(err)
	}

	url := app.CreatedUrl(*flags, *defVal)
	resUrl, err := app.GetUrl(url)
	if err != nil {
		log.Fatal(err)
	}

	fotmat, err := app.GetFormat(resUrl)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(flags.Name + fotmat)
	if err != nil {
		log.Fatalf("unable to create file %v", err)
	}

	_, err = io.Copy(file, resUrl.Body)
	if err != nil {
		log.Fatalf("unable to write file %v", err)
	}

}
