package main

import "fmt"

func main()  {
	arr := []int{1,2,3,4,5,6,7,8,9,10}

	for _,num := range arr{
		if num%2==0 {
			fmt.Println(num," is even number")
		} else {
			fmt.Println(num,"is odd number")
		}
	}
}
