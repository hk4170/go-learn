package main

import (
	"fmt"
	"log"

	"go-learn/lib"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.

	//log.SetPrefix("error demo: ")//设置前缀
	//log.SetFlags(0) //不显示时间

	// Request a greeting message.
	message, err := lib.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
