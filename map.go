package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	//dict := make(map[string]string)
	dict := map[string]string{}
	dict["name"] = "jnoh"
	dict["age"] = "20"
	fmt.Printf("hello %s your age is %s",dict["name"],dict["age"])
	file,_ := os.Create("map.gob")
	encoding := gob.NewEncoder(file)
	encoding.Encode(dict)

	println()
	type Person struct {
		Info map[string]string
	}
	p := Person{
		Info: map[string]string{
			"name": "Alice",
			"city": "Beijing",
		},}
	fmt.Println(p.Info["name"]) // 输出: Alice		
    
}