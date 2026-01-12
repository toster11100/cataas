package app

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"main.go/internal/config"
)

type Cat struct {
	url     *url.URL
	resBody []byte
	name    string
	conType string
}

func New(config config.Config) *Cat {
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

	cat := &Cat{
		url:  catURL,
		name: config.Name,
	}
	return cat
}

func (cat *Cat) GetCat() error {
	if err := cat.getRes(); err != nil {
		return err
	}

	if err := cat.getFormat(); err != nil {
		return err
	}
	return nil
}

func (cat *Cat) getRes() error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	cnt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(cnt, http.MethodGet, cat.url.String(), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("website access problems: %v", err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d", res.StatusCode)
	}

	cat.conType = res.Header.Get("Content-Type")
	if cat.conType == "" {
		return errors.New("empty Content-Type header")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	cat.resBody = body
	return nil
}

func (cat *Cat) getFormat() error {
	var format string
	switch cat.conType {
	case "image/png":
		format = ".png"
	case "image/jpeg":
		format = ".jpeg"
	default:
		return fmt.Errorf("unknown format: %s", cat.conType)
	}

	cat.name += format
	return nil
}

func (cat *Cat) SavePicture() error {
	file, err := os.Create(cat.name)
	if err != nil {
		return fmt.Errorf("unable to create file %v", err)
	}

	_, err = io.Copy(file, bytes.NewReader(cat.resBody))
	if err != nil {
		return fmt.Errorf("unable to write file %v", err)
	}

	return file.Close()
}
