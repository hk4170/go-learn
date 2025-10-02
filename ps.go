package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-ps"
)

type proce struct{
    pid int 
	exec string
}

var process []proce

func main() {
	processes, err := ps.Processes()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range processes {
		//fmt.Println(p)
		var ps proce
		ps.pid = p.Pid()
		ps.exec = p.Executable()
		process = append(process, ps)

		//fmt.Printf("PID: %d, Name: %s, PPID: %d\n", p.Pid(), p.Executable(), p.PPid())
	}
	for _,ps := range process{
		if ps.exec == "go"{
			fmt.Println(ps.pid)
		}
		
	}
}
