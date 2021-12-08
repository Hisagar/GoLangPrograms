package main

import "fmt"

func main()  {

	 arr:= [5]int {1,2,3,4,5}

	for i:= 0; i< len(arr); i++ {
		fmt.Println(arr[i])
	}

	var names [3]string
	names[0] = "Sagar"
	names[1] = "Rohit"
	names[2] = "Jay"
	 fmt.Println(names)


}