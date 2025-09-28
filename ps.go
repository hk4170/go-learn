package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-ps"
)

func main() {
	processes, err := ps.Processes()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range processes {
		fmt.Printf("PID: %d, Name: %s, PPID: %d\n", p.Pid(), p.Executable(), p.PPid())
	}
}
