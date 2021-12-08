package main

import "fmt"

type person struct {
	firstname string
	lastname string
}
func main()  {
	sagar := person{firstname: "Sagar", lastname: "Wankhade"}
	fmt.Println(sagar)
	var user person
	fmt.Println(user)
	fmt.Printf("%+v\n", user)

	var saurav person
	saurav.firstname = "Saurav"
	saurav.lastname = "Wankhade"
	fmt.Printf("%+v", saurav)
}
