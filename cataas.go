package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Tag    string `yaml:"tag"`
	Says   string `yaml:"says"`
	Filter string `yaml:"filter"`
	Height int    `yaml:"height"`
	Widht  int    `yaml:"widht"`
}

type Flag struct {
	Tag, Says, Filter, Config *string
	Height, Width             *int
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please give me one argument!\n")
	}
	arg := Flag{}
	arg.Tag = pflag.StringP("tag", "t", "", "tag cats")
	arg.Says = pflag.StringP("says", "s", "", "cat will say hello")
	arg.Filter = pflag.StringP("filter", "f", "", "filter for cute cats")
	arg.Height = pflag.IntP("height", "h", 0, "image height")
	arg.Config = pflag.StringP("config", "c", "./config.yaml", "yaml config")
	arg.Width = pflag.IntP("width", "w", 0, "image width")
	pflag.Parse()

	filename := pflag.Arg(0)
	if filename == "" {
		log.Fatal("no name in arguments\n")
	}

	FileYAML, err := os.ReadFile(*arg.Config)
	if err != nil {
		log.Fatal("problem with yaml config\n", err)
	}
	c := YamlConfig{}
	yaml.Unmarshal(FileYAML, &c)

	URL := c.SayMyURL(arg)

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

func (c YamlConfig) SayMyURL(arg Flag) url.URL {
	if *arg.Tag == "" {
		*arg.Tag = c.Tag
	}
	if *arg.Says == "" {
		*arg.Says = c.Says
	}
	if *arg.Filter == "" {
		*arg.Filter = c.Filter
	}
	if *arg.Height == 0 {
		*arg.Height = c.Height
	}
	if *arg.Width == 0 {
		*arg.Width = c.Widht
	}
	v := url.Values{}
	if *arg.Filter != "" {
		v.Set("filter", *arg.Filter)
	}
	if *arg.Width != 0 {
		v.Set("width", strconv.Itoa(*arg.Width))
	}
	if *arg.Height != 0 {
		v.Set("height", strconv.Itoa(*arg.Height))
	}
	URL := url.URL{
		Scheme:   "https",
		Host:     "cataas.com",
		Path:     "/cat",
		RawQuery: v.Encode(),
	}
	URL = *URL.JoinPath(*arg.Tag)
	if *arg.Says != "" {
		URL = *URL.JoinPath("says", *arg.Says)
	}
	return URL
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
