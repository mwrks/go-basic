package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Number", i)
	// }

	i := 0
	// for i < 5 {
	// 	fmt.Println("Number", i)
	// 	i++
	// }

	for {
		fmt.Println("Number", i)
		i++
		if i == 5 {
			break
		}
	}
}
