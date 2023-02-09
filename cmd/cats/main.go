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

	file, err := os.Create(config.Name + cataas.Format)
	if err != nil {
		log.Fatalf("unable to create file %v", err)
	}

	_, err = io.Copy(file, cataas.Res.Body)
	if err != nil {
		log.Fatalf("unable to write file %v", err)
	}

	defer file.Close()
	defer cataas.Res.Body.Close()
}
