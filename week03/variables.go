// variables are used to store info to be referenced and manipulated in a computer program. they also provide a way of labeling data with a descriptive name, so our programs can be understood more clearly by the treader and ourselves. it is helpful to think of variables as containers that hold information. their sole purpose is ot label and store data in memory. this data can then be used throughout your program.

package main

import "fmt" //formatted I/O

func main(){
	// explicit declaration followed by assignment
	// var = keyword, name = variable name, string = data type
	var name0 string
	name0 = "Musme"

	//explicit declaration and assignment
	var name1 string = "Eunice"

	// implicit declaration and assignment
	name2 := "Bravo"

	// first variable with mastermnd lol
	var age int
	age = 25

	fmt.Println("name:", name0, name1, name2, "age:", age)
}