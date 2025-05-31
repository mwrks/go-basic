package main

import "fmt"

func main() {
	var grade = 8
	if grade == 10 {
		fmt.Println("Passed with excellent grade")
	} else if grade > 5 {
		fmt.Println("Passed")
	} else if grade == 4 {
		fmt.Println("Passed with minimum grade")
	} else {
		fmt.Printf("Failed with grade %d\n", grade)
	}
}
