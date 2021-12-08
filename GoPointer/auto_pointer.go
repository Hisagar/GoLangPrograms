package main

import "fmt"

type person struct {
	fname string
	lname string
}


func main()  {
	sagar := person{
		fname: "Sagar",
		lname: "Wankhade",
	}

	sagar.updateNameWithPointer("Saurav")
	sagar.print()
}

func (pointerToPerson *person) updateNameWithPointer(newFname string)  {
	(*pointerToPerson).fname = newFname
}

func (p person) print()  {
	fmt.Printf("%+v", p)
}


