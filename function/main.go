package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"John", "Wick"}
	printMessage("Hello", names)

}

func printMessage(message string, arr []string) {
	// nameString := strings.Join(arr, " ")
	// fmt.Printf("Message : \"%s %s\"", message, nameString)
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
	
}
