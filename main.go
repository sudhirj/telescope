package main

import (
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/sudhirj/telescope/telescope"
)

var port = flag.String("port", ":8353", "Port to listen on. Defaults to :8353")
var host = flag.String("host", "", "Host to fetch images from. If this option isn't specified, Telescope assumes the first part of the path is the base64 encoded host")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	http.Handle("/", http.HandlerFunc(serve))
	http.HandleFunc("/favicon.ico", func(r http.ResponseWriter, req *http.Request) {})

	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}

func chooseURL(req *http.Request) (*url.URL, error) {
	opts := telescope.NewOptionsFromRequest(req)
	if *host != "" {
		imageURL, err := url.Parse(*host)
		if err != nil {
			return nil, err
		}
		imageURL.Path = req.URL.Path
		return imageURL, nil
	} else {
		return opts.SourceURL()
	}
}

func serve(rw http.ResponseWriter, req *http.Request) {
	opts := telescope.NewOptionsFromRequest(req)
	imageURL, err := chooseURL(req)

	if err != nil {
		rw.WriteHeader(404)
		return
	}

	imageRequest, err := http.Get(imageURL.String())
	if err != nil {
		rw.WriteHeader(404)
		return
	}
	defer imageRequest.Body.Close()

	sourceImage, format, err := image.Decode(imageRequest.Body)
	if err != nil {
		rw.WriteHeader(400)
		return
	}

	rw.Header().Add("Cache-Control", "public, max-age=864000")
	rw.Header().Add("Content-Type", strings.Join([]string{"image", format}, "/"))

	wipImage := sourceImage

	if opts.Width > 0 || opts.Height > 0 {
		wipImage = imaging.Resize(sourceImage, int(opts.Width), int(opts.Height), imaging.Lanczos)
	}

	if opts.Blur > 0 {
		wipImage = imaging.Blur(wipImage, opts.Blur)
	}

	if format == "jpeg" {
		jpeg.Encode(rw, wipImage, nil)
	}

	if format == "png" {
		png.Encode(rw, wipImage)
	}

	go runtime.GC()
}
