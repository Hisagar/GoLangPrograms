package main

import "fmt"

type contactInfo struct {
	email string
	pincode int
}
type personInfo struct {
	firstname string
	lastname string
	contact contactInfo
}

func main()  {
	sagar := personInfo{
		firstname: "Sagar",
		lastname: "Wankhade",
		contact: contactInfo{
			email: "sagar@example.com",
			pincode: 394210,
		},
	}
	psagar := &sagar
	psagar.updateNameWithPointer("Saurav")
	sagar.print()
}

func (pointerToPerson *personInfo) updateNameWithPointer(newFname string)  {
	(*pointerToPerson).firstname = newFname
}

func (p personInfo) print()  {
	fmt.Printf("%+v", p)
}


