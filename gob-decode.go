package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
    type host  struct{
        Host string
        Port int
        proxy string
        Proxy string
        Os string
    }
    var hosts []host

    file ,err := os.Open("test.data")
    if err != nil{
        log.Fatal(err)
    }
    defer file.Close()
    
    dec := gob.NewDecoder(file)
    dec.Decode(&hosts)
    

    fmt.Printf("Decoded: %+v\n", hosts)
}
