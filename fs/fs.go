package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	var dir string
	port := flag.String("port", "5151", "port to serve HTTP on")
	path := flag.String("path", "", "path to serve")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}