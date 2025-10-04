package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fruits := []string{
		"苹果", 
		"香蕉", 
		"橘子", 
		"葡萄", 
		"西瓜",
	}
	rand.Shuffle(len(fruits), func(i, j int) {
		fruits[i], fruits[j] = fruits[j], fruits[i]
	})
	fmt.Println(fruits) // 打乱顺序每次运行顺序可能不同
	for range 10{
	    fmt.Println(fruits[rand.Intn(len(fruits))])
	}
}
