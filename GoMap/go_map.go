package main

import "fmt"
func main()  {
	var saurav map[string]string
	testmap := make(map[string]string)
	testmap["fname"]="abc"
	testmap["lname"]="xyz"
	sagar := map[string]string{
		"fname":"Sagar",
		"lname":"Wankhade",
	}
	delete(sagar ,"fname")
	fmt.Println(sagar)
	fmt.Println(saurav)
	fmt.Println(testmap)
}
