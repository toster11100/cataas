package app

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"main.go/interal/config"
)

type Cat struct {
	URL    *url.URL
	Res    *http.Response
	Format string
}

func New(config config.Config) (*Cat, error) {
	cat := &Cat{}

	cat.createURL(config)
	if err := cat.getRes(); err != nil {
		return nil, err
	}
	if err := cat.getFormat(); err != nil {
		return nil, err
	}

	return cat, nil
}

func (cat *Cat) createURL(config config.Config) {
	v := url.Values{}

	if config.Filter != "" {
		v.Set("filter", config.Filter)
	}
	if config.Width != 0 {
		v.Set("width", strconv.Itoa(config.Width))
	}
	if config.Height != 0 {
		v.Set("height", strconv.Itoa(config.Height))
	}

	catURL := &url.URL{
		Scheme:   "https",
		Host:     "cataas.com",
		Path:     "/cat",
		RawQuery: v.Encode(),
	}

	catURL = catURL.JoinPath(config.Tag)
	if config.Say != "" {
		catURL = catURL.JoinPath("says", config.Say)
	}
	cat.URL = catURL
}

func (cat *Cat) getRes() error {
	res, err := http.Get(cat.URL.String())
	if err != nil {
		return fmt.Errorf("website access problems: %v", err)
	}
	cat.Res = res

	if cat.Res.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d", cat.Res.StatusCode)
	}

	return nil
}

func (cat *Cat) getFormat() error {
	contentType := cat.Res.Header.Get("Content-Type")
	if contentType == "" {
		return errors.New("empty Content-Type header")
	}

	format := ""
	switch contentType {
	case "image/png":
		format = ".png"
	case "image/jpeg":
		format = ".jpeg"
	default:
		return fmt.Errorf("unknown format: %s", contentType)
	}
	cat.Format = format
	return nil
}
