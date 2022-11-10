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
	Tag    *string
	Says   *string
	Filter *string
	Config *string
	Height *int
	Width  *int
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please give me one argument!")
	}
	arg := Flag{}
	arg.Tag = pflag.StringP("tag", "t", "", "tag cats")
	arg.Says = pflag.StringP("says", "s", "", "cat will say hello")
	arg.Filter = pflag.StringP("filter", "f", "", "filter for cute cats")
	arg.Config = pflag.StringP("config", "c", "./config.yaml", "yaml config")
	arg.Height = pflag.IntP("height", "h", 0, "image height")
	arg.Width = pflag.IntP("width", "w", 0, "image width")
	pflag.Parse()

	filename := pflag.Arg(0)
	if filename == "" {
		log.Fatal("no name in arguments")
	}

	fileYAML, err := os.ReadFile(*arg.Config)
	if err != nil {
		log.Fatalf("problem with yaml confi %v", err)
	}
	c := YamlConfig{}
	yaml.Unmarshal(fileYAML, &c)

	URL := c.BuildURL(arg)

	response, err := http.Get(URL.String())
	if err != nil {
		log.Fatalf("website access problems %v", err)
	}
	defer response.Body.Close()
	TypeImage, err := GetFormat(response.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}

	if filename == "-" {
		if _, err := io.Copy(os.Stdout, response.Body); err != nil {
			log.Fatalf("cannot be output to command line %v", err)
		}

	} else {
		file, err := os.Create(filename + TypeImage)
		if err != nil {
			log.Fatalf("unable to create file %v", err)
		}

		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatalf("unable to write file %v", err)
		}
	}
}

func (c YamlConfig) BuildURL(arg Flag) *url.URL {
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
	URL := &url.URL{
		Scheme:   "https",
		Host:     "cataas.com",
		Path:     "/cat",
		RawQuery: v.Encode(),
	}
	URL = URL.JoinPath(*arg.Tag)
	if *arg.Says != "" {
		URL = URL.JoinPath("says", *arg.Says)
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
