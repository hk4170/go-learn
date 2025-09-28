package main

import (
	"fmt"
	"os/exec"
	
)

func main() {
	cmd := exec.Command("service","mysql","start")
	cmd.Run()
	fmt.Println(cmd.Output())
    fmt.Println()
}
