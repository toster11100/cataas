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
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please give me one argument!")
	}
	var Tag *string = pflag.StringP("tag", "t", "", "tag cats")
	var Says *string = pflag.StringP("says", "s", "", "cat will say hello")
	var Filter *string = pflag.StringP("filter", "f", "", "filter for cute cats") //blur, mono, sepia, negative, paint, pixel
	var Height *int = pflag.IntP("height", "h", 0, "image height")
	var Width *int = pflag.IntP("width", "w", 0, "image width")
	pflag.Parse()

	filename := pflag.Arg(0)
	if filename == "" {
		log.Fatal("no name in arguments")
	}

	URL := SayMyURL(*Height, *Width, *Tag, *Filter, *Says)
	response, err := http.Get(URL.String())
	if err != nil {
		log.Fatal("website access problems", err)
	}
	defer response.Body.Close()
	TypeImage, err := GetFormat(response.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}

	if filename == "-" {
		if _, err := io.Copy(os.Stdout, response.Body); err != nil {
			log.Fatal("cannot be output to command line", err)
		}

	} else {
		file, err := os.Create(filename + TypeImage)
		if err != nil {
			log.Fatal("unable to create file\n", err)
		}

		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal("unable to write file", err)
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
