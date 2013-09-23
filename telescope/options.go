package telescope

import (
	"net/http"
	"strconv"
)

type Options struct {
	Origin string
	Width  uint
	Height uint
}

func NewOptionsFromMap(options map[string]string) *Options {
	p := new(Options)
	p.Origin = options["origin"]

	parse := func(opt string) (uint64, error) {
		return strconv.ParseUint(options[opt], 10, 0)
	}

	width, err := parse("width")
	if err == nil {
		p.Width = uint(width)
	}

	height, err := parse("height")
	if err == nil {
		p.Height = uint(height)
	}

	return p
}

func NewOptionsFromRequest(req *http.Request) *Options {
	return NewOptionsFromMap(map[string]string{
		"width":  req.FormValue("w"),
		"height": req.FormValue("h"),
		"origin": req.FormValue("origin"),
	})
}
