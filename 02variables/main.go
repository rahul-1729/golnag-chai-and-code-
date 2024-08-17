package main

import "fmt"

func main() {
	var username string = "Rahul"
	fmt.Println(username)
	fmt.Printf("Variable ids of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable ids of type: %T \n", isLoggedIn)

	var smallval int = 256
	fmt.Println(smallval)
	fmt.Printf("Variable ids of type: %T \n", smallval)

	var smallfloat float64 = 255.88888888888888
	fmt.Println(smallfloat)
	fmt.Printf("Variable ids of type: %T \n", smallfloat)

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("variable is of type: %T \n", anotherVar)

	var website = "learncodeonline.in"
	fmt.Println(website)

	numberOfUser := 30000 //val
	fmt.Println(numberOfUser)
}
