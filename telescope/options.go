package telescope

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Options struct {
	Origin string
	Width  uint
	Height uint
	Blur   float64
	File   string
}

func (opts *Options) SourceURL() (*url.URL, error) {
	baseUrl, err := url.Parse(opts.Origin)
	if err != nil {
		return nil, err
	}
	baseUrl.Path = opts.File
	return baseUrl, nil
}

func NewOptionsFromRequest(req *http.Request) *Options {

	opts := new(Options)

	parseInt := func(opt string) (uint64, error) {
		return strconv.ParseUint(opt, 10, 0)
	}

	parseFloat := func(opt string) (float64, error) {
		return strconv.ParseFloat(opt, 64)
	}

	width, err := parseInt(req.FormValue("w"))
	if err == nil {
		opts.Width = uint(width)
	}

	height, err := parseInt(req.FormValue("h"))
	if err == nil {
		opts.Height = uint(height)
	}

	blur, err := parseFloat(req.FormValue("blur"))
	if err == nil {
		opts.Blur = blur
	}

	pathParts := strings.SplitN(req.URL.Path, "/", 3)

	origin, err := base64.URLEncoding.DecodeString(pathParts[1])
	if err == nil {
		opts.Origin = string(origin)
	}

	opts.File = pathParts[2]

	return opts
}
