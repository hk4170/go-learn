package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)
type host struct {
    Host string
    Os string
    Os_version string
    Port int
    Proxy string 
    proxy string
}
var hosts []host

func main() {
    if len(os.Args) != 2{
        log.Fatalf("gob-decode <file>")
    }

    file ,err := os.Open(os.Args[1])
    if err != nil{
        log.Fatal(err)
    }
    defer file.Close()

    dec := gob.NewDecoder(file)
    dec.Decode(&hosts)
    
    
    fmt.Printf("Decoded: %+v\n", hosts)
}
