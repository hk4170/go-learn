package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed web
var webFS embed.FS

func main() {
	fsys, _ := fs.Sub(webFS, "web")
	http.Handle("/", http.FileServer(http.FS(fsys)))

	http.ListenAndServe(":8080", nil)
}

