package app

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"main.go/interal/config"
)

type Cat struct {
	URL     *url.URL
	Res     *http.Response
	Format  string
	catConf config.Config
}

func New(config config.Config) (*Cat, error) {
	cat := &Cat{}
	cat.catConf = config

	cat.createURL()
	if err := cat.getRes(); err != nil {
		return nil, err
	}
	if err := cat.getFormat(); err != nil {
		return nil, err
	}

	return cat, nil
}

func (cat *Cat) createURL() {
	v := url.Values{}

	if cat.catConf.Filter != "" {
		v.Set("filter", cat.catConf.Filter)
	}
	if cat.catConf.Width != 0 {
		v.Set("width", strconv.Itoa(cat.catConf.Width))
	}
	if cat.catConf.Height != 0 {
		v.Set("height", strconv.Itoa(cat.catConf.Height))
	}

	catURL := &url.URL{
		Scheme:   "https",
		Host:     "cataas.com",
		Path:     "/cat",
		RawQuery: v.Encode(),
	}

	catURL = catURL.JoinPath(cat.catConf.Tag)
	if cat.catConf.Say != "" {
		catURL = catURL.JoinPath("says", cat.catConf.Say)
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

func (cat *Cat) SavePicture() error {
	file, err := os.Create(cat.catConf.Name + cat.Format)
	if err != nil {
		return fmt.Errorf("unable to create file %v", err)
	}

	_, err = io.Copy(file, cat.Res.Body)
	if err != nil {
		return fmt.Errorf("unable to write file %v", err)
	}

	defer file.Close()
	return nil
}
