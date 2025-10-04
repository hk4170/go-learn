package main

import (
	"database/sql"
	"fmt"
	"log"
    _ "github.com/mattn/go-sqlite3"
)

func main(){
	db ,err := sql.Open("sqlite3","test.db")
	//db ,err := sql.Open("sqlite3","test.db")
    if err != nil {
		log.Fatal(err)
	}
	rows ,err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	type person struct{
		Id int
		Name string
		Age int
	}
	var users []person
	for rows.Next(){
		var u person
		rows.Scan(&u.Id,&u.Name,&u.Age)
		users = append(users, u)
	}
	
	for _,v := range users{
       fmt.Printf("name: %s,Age: %v \n",v.Name,v.Age) 
	}
	

    


}