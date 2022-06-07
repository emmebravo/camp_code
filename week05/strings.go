package main

import "fmt"

func main(){
	var firstName string
	firstName = "Musme"

	surname := "Bravo"

	fullName := fmt.Sprintf(firstName + " " + surname)

	// fmt.Println(firstName)
	// fmt.Println(surname)
	fmt.Println(fullName)
	fmt.Printf("Hello my name is %v %v\n", firstName, surname)
}