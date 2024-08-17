package main

import "fmt"

const LoginToken string = "ghabbhhjd" //public

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

	numberOfUser := 30000 //valras operator is allowed only insde a method or function
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)

}
