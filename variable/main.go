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

	var integer int = 56
	fmt.Println(integer)

	var float float32 = 17.56423
	fmt.Printf("The number is: %.3f\n", float)

	var truth bool = true
	fmt.Printf("Q:\tAm I a programmer?\nA:\t%t\n", truth)

	message := `Today
Hope I can keep my work
Hope I can succeed`
	fmt.Println(message)
	
}
