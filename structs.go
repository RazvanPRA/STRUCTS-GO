package main

import (
	"fmt"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	age := getUserData("Please enter your birthdate (YYYY-MM-DD): ")

	// ... do something awesome with that gathered data!

	fmt.Println(firstName, lastName, age)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}