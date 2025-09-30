package main

import "fmt"

func main(){
    type person struct{
		name string
		age int
		email string
	}
	var users []person
	var p person
	p.name = "sjsjs"
	p.age = 25
	p.email = "sjskdn@12.89"
	users = append(users, p)
	users = append(users, p)
	users = append(users, p)
	fmt.Println(users)
	fmt.Println(users[0])
}