package main

import "fmt"
func main()  {

	fmt.Println("Simple Loop")
	for i:=1; i<=10;i++{
		fmt.Println(i)
	}

	fmt.Println("While Loop")
	j:=1;
	for ;j<=10;{
		fmt.Println(j)
		j++
	}

	// Do while loop
	fmt.Println("Do While Loop")
	k:=1
	for{
		fmt.Println(k)
		k++
		if k>10{
			break
		}
	}

	fmt.Println("Fullname : "+concatName("Sagar","Wankhade"))



}

func concatName(fname ,  lname string) string {
	return fname+" "+lname
}
