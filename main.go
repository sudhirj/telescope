package main

import "net/http"
import "fmt"

import telescope "github.com/sudhirj/telescope/telescope"

func fetch(w http.ResponseWriter, req *http.Request) {
	options := telescope.NewOptionsFromRequest(req)
	picture := telescope.LoadPicture(options.Origin)
	defer picture.Destroy()
	picture.SizeToWidth(options.Width)
	picture.WriteTo(w)
}

func favicon(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
}

func main() {
	fmt.Println("Ready.")
	telescope.Start()
	defer telescope.Stop()
	http.HandleFunc("/", fetch)
	http.HandleFunc("/favicon.ico", favicon)
	http.ListenAndServe(":9000", nil)
}
