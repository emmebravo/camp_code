package main

import "fmt"

func main() {
	name := "musme"

	fmt.Println(name[0]) //will print rune not string
	fmt.Println(string(name[0])) //will print string

	for i := 0; i < len(name); i++ {
		fmt.Println("The next letter in the string is: ", string(name[i]))
	}
}