package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("service","mysql","start")
	//exec.Command("service","mysql","start").Run()
	cmd.Run()
	fmt.Println(cmd.Output())
}
