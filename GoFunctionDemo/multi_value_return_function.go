package main

import "fmt"
func main()  {
	str, num := multi_value_return_function()
	fmt.Println("String value : ",str+"\n"+"Int Value :",num)
}

func multi_value_return_function()(string, int)  {
	return "xyz" , 1
}