package main

import "fmt"

func main() {
	var firstName string = "John"
	var lastName string
	lastName = "Wick"
	fmt.Printf("Halo %s %s!\n", firstName, lastName)

	person := "Saya"
	fmt.Printf("%s adalah programmer!\n", person)

	_ = "unused"
	_, secondVar := "ignored", "used"
	fmt.Println(secondVar)

	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)
}
