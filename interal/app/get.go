package app

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"main.go/interal/config"
	"main.go/interal/flags"
)

func CreatedUrl(flag flags.Flags, defFlag config.DefaultFlags) *url.URL {
	if flag.Tag == "" {
		flag.Tag = defFlag.Tag
	}
	if flag.Says == "" {
		flag.Says = defFlag.Says
	}
	if flag.Filter == "" {
		flag.Filter = defFlag.Filter
	}
	if flag.Height == 0 {
		flag.Height = defFlag.Height
	}
	if flag.Width == 0 {
		flag.Width = defFlag.Width
	}
	v := url.Values{}

	if flag.Filter != "" {
		v.Set("filter", flag.Filter)
	}
	if flag.Width != 0 {
		v.Set("width", strconv.Itoa(flag.Width))
	}
	if flag.Height != 0 {
		v.Set("height", strconv.Itoa(flag.Height))
	}

	u := &url.URL{
		Scheme:   "https",
		Host:     "cataas.com",
		Path:     "/cat",
		RawQuery: v.Encode(),
	}

	u = u.JoinPath(flag.Tag)
	if flag.Says != "" {
		u = u.JoinPath("says", flag.Says)
	}

	return u
}

func GetUrl(url *url.URL) (*http.Response, error) {
	resUrl, err := http.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("website access problems %v", err)
	}
	defer resUrl.Body.Close()

	return resUrl, nil
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
