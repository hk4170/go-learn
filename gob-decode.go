package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
    if len(os.Args) != 2{
        log.Fatalf("gob-decode <file>")
    }
    type host  struct{
        Host string
        Port int
        proxy string
        Proxy string
        Os string
    }
    var hosts []host

    file ,err := os.Open(os.Args[1])
    if err != nil{
        log.Fatal(err)
    }
    defer file.Close()

    dec := gob.NewDecoder(file)
    dec.Decode(&hosts)
    
    
    fmt.Printf("Decoded: %+v\n", hosts)
}
