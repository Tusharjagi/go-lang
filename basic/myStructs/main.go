package main

import "fmt"

func main () {
	fmt.Println("Welcome to struct")
	// no inheritance in golang; No super or parent
	tushar := User{"Tushar", "tusharjagi@gmail.com", true, 18}

	fmt.Println(tushar)
	fmt.Printf("tushar details are %+v\n", tushar)
	fmt.Printf("Name is %v and email is %v\n", tushar.Name, tushar.Email)

}

type User struct {
	Name 		string 
	Email 	string
	Status	bool
	Age 		int
}