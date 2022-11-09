package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

type YamlConfigs struct {
	Tag    string `yaml:"tag"`
	Says   string `yaml:"says"`
	Filter string `yaml:"filter"`
	Height int    `yaml:"height"`
	Widht  int    `yaml:"widht"`
}

var (
	Tag, Says, Filter, Config *string
	Height, Width             *int
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please give me one argument!\n")
	}

	Tag = pflag.StringP("tag", "t", "", "tag cats")
	Says = pflag.StringP("says", "s", "", "cat will say hello")
	Filter = pflag.StringP("filter", "f", "", "filter for cute cats")
	Height = pflag.IntP("height", "h", 0, "image height")
	Width = pflag.IntP("width", "w", 0, "image width")
	Config = pflag.StringP("config", "c", "./config.yaml", "yaml config")
	pflag.Parse()

	filename := pflag.Arg(0)
	if filename == "" {
		log.Fatal("no name in arguments\n")
	}

	FlagOrYaml(*Height, *Width, *Tag, *Filter, *Says, *Config)

	URL := SayMyURL(*Height, *Width, *Tag, *Filter, *Says)
	response, err := http.Get(URL.String())
	if err != nil {
		log.Fatal("website access problems\n", err)
	}
	defer response.Body.Close()
	TypeImage, err := GetFormat(response.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}

	if filename == "-" {
		if _, err := io.Copy(os.Stdout, response.Body); err != nil {
			log.Fatal("cannot be output to command line\n", err)
		}

	} else {
		file, err := os.Create(filename + TypeImage)
		if err != nil {
			log.Fatal("unable to create file\n", err)
		}

		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal("unable to write file\n", err)
		}
	}
}

func SayMyURL(height, width int, tag, filter, says string) url.URL {
	v := url.Values{}
	if filter != "" {
		v.Set("filter", filter)
	}
	if width != 0 {
		v.Set("width", strconv.Itoa(width))
	}
	if height != 0 {
		v.Set("height", strconv.Itoa(height))
	}
	url := url.URL{
		Scheme:   "https",
		Host:     "cataas.com",
		Path:     "/cat",
		RawQuery: v.Encode(),
	}
	url = *url.JoinPath(tag)
	if says != "" {
		url = *url.JoinPath("says", says)
	}
	fmt.Println(&url)
	return url
}

func GetFormat(format string) (string, error) {
	switch {
	case format == "image/png":
		return ".png", nil
	case format == "image/jpeg":
		return ".jpeg", nil
	}
	return "", errors.New("unknow format")
}

func FlagOrYaml(height, width int, tag, filter, says, config string) {
	c := YamlConfigs{}
	FileYAML, err := os.ReadFile(*Config)
	if err != nil {
		log.Fatal("problem with yaml config\n", err)
	}
	err = yaml.Unmarshal(FileYAML, &c)
	if err != nil {
		log.Fatal("problem with unmarshal\n", err)
	}
	if *Tag == "" {
		Tag = &c.Tag
	}
	if *Says == "" {
		Says = &c.Says
	}
	if *Filter == "" {
		Filter = &c.Filter
	}
	if *Height == 0 {
		Height = &c.Height
	}
	if *Width == 0 {
		Width = &c.Widht
	}
}
