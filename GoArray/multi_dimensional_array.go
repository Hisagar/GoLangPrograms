package main

import "fmt"
func main()  {

	colors:= [3][3]string{
		{"red", "green", "blue"},
		{"orange", "pink", "white"},
		{"black", "gray", "yellow"},
	}

	for x:= 0; x<=2; x++ {
		for y:=0; y<=2; y++ {
			fmt.Println(colors[x][y])
		}
	}
}
