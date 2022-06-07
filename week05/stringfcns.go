package main

import (
	"fmt"
	"strings"
)

func main(){
	// imagine we are collecting email addresses for a mailing list
	// for email, case does not matter. normalize this by making all lowercase.
	email := "CoolPerson.Billy@gmail.com"

	normalized_email := strings.ToLower(email)
	upper_email := strings.ToUpper(email)

	fmt.Println(normalized_email)
	fmt.Println(upper_email)

	// imagine you are collecting street addresses to go into a database.
	// let's split the information to make it easier to organize it all
	address := "1234 Fancy House Lane, Pittsburgh, PA., 20895"

	processed_address := strings.Split(address, ",")

	fmt.Printf("%#v\n", processed_address)

	// check if November 20 is the person's birthday, if so they win a prize
	birthday := "November 20"

	birthdayTwin := strings.Contains(birthday, "November 20")

	fmt.Println(birthdayTwin)

	// using repeat to repeat lyrics
	lyrics := "baby"

	ashanti_lyrics := strings.Repeat(lyrics, 20)

	fmt.Println(ashanti_lyrics)

	// using repeat to create a sections
	fmt.Println(strings.Repeat("*", 50))
	fmt.Println("instructions")
	fmt.Println(strings.Repeat("*", 50))

	//
}