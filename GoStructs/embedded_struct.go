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

	sagar.updateName("Saurav")
	sagar.print()
}

func (p personInfo) updateName(newFname string)  {
	p.firstname = newFname
}

func (p personInfo) print()  {
	fmt.Printf("%+v", p)
}
