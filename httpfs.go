package frontend
import (
	."ashe/core"
	"net/http"
	"embed"
	"io/fs"
)

//go:embed web/*
var webFS embed.FS
func FrontendServer(){
	if Config.Server.DisFrontend {
		return
	}
	web , _ := fs.Sub(webFS,"web")
	http.Handle("/", http.FileServer(http.FS(web)))
	http.ListenAndServe(":"+Config.Server.FrontendPort,nil)
}