//go:build windows
package main

import (
	"github.com/pkg/browser"
)

func open(){
	browser.OpenURL("http://localhost:8080")
}

func main(){
	open()
}