package main

import "fmt"
func main()  {
	name := map[string]string{
		"fname":"Sagar",
		"lname":"Wankhade",
	}
	for n, key := range name{
		fmt.Println(n)
		fmt.Println(key)
	}

}
