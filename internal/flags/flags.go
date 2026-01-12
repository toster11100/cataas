package flags

import (
	"errors"
	"os"

	"github.com/spf13/pflag"
)

type Flags struct {
	Name   string
	Tag    string
	Say    string
	Filter string
	Config string
	Height int
	Width  int
}

func ParseOptions() (*Flags, error) {
	if len(os.Args) == 1 {
		return nil, errors.New("please provide at least one argument")
	}

	flags := &Flags{}
	pflag.StringVarP(&flags.Tag, "tag", "t", "", "tag cats")
	pflag.StringVarP(&flags.Say, "say", "s", "", "text for image")
	pflag.StringVarP(&flags.Filter, "filter", "f", "", "filter for image")
	pflag.StringVarP(&flags.Config, "config", "c", "./config.yaml", "yaml config")
	pflag.IntVarP(&flags.Height, "height", "h", 0, "image height")
	pflag.IntVarP(&flags.Width, "width", "w", 0, "image width")
	pflag.Parse()

	flags.Name = pflag.Arg(0)
	if flags.Name == "" {
		return nil, errors.New("no filename provided in the arguments")
	}

	if pflag.NArg() > 1 {
		return nil, errors.New("only one non-flag argument is allowed")
	}

	return flags, nil
}
