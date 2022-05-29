// data values that stay the same every time a program is executed

package main

import "fmt"

func main(){
	//my first constant
	const age int = 26

	const age1 = 43
	// unlike variable, you can't write const age int & then reassign age on the line below it

	fmt.Println(age, age1) //this print adds a ln break at the end of the output
}