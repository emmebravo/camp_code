package main

import "fmt"

func init() {
	fmt.Println("Welcome to init() fcn")
}

func init(){
	fmt.Println("Hello! init() fcn")
}

func main() {
	fmt.Println("Welcome to main() fcn")
}