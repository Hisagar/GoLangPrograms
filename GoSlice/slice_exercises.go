package main

import "fmt"

func main() {

	slice := []string{"a", "b", "c", "d",
		"e", "f", "g"}

	fmt.Println("Array:", slice)
	slice1 := slice[1:6]
	slice2 := slice[2:4]
	slice3 := slice[0:7]
	fmt.Println("Slice:", slice1)
	fmt.Println("Slice:", slice2)
	fmt.Println("Slice:", slice3)
}