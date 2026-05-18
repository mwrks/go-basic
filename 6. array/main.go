package main

import "fmt"

func main() {
	names := [4]string{}
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	// names[4] = "is trash" //this line causes error due to unallocated array index
	fmt.Println(names[0], names[1], names[2], names[3])

	// fruits := [4]string{"apple", "grape", "banana", "melon"} //horizontal initialization
	fruits := [4]string{
		"apple",
		"grape",
		"banana",
		"melon", // (add comma at the end of the last index)
	} //vertical initialization

	fmt.Println("Sum of elements \t:", len(fruits))
	fmt.Println("Array values \t\t:", fruits)

	numbers := [...]int{2, 3, 2, 4, 3}

	fmt.Println("Array values \t\t:", numbers)
	fmt.Println("Jumlah elemen \t\t:", len(numbers))

	numbers1 := [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}} //array type redundant
	numbers2 := [2][3]int{{3, 2, 3}, {3, 4, 5}}             //preferred way
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// fruits := [4]string{"apple", "grape", "banana", "melon"} //ommited since it's already declared on top
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Element %d : %s\n", i, fruits[i])
	}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// for i, fruit := range fruits {
	// 	fmt.Printf("nama buah : %s\n", fruit)
	// } //this code throws an error since there's an unused variable, thus...

	for _, fruit := range fruits { //use the underscore variable to ommit the value of the first function return variable of the range function
		fmt.Printf("Fruit : %s\n", fruit)
	}

	//if you only need the first function return variable you can do these
	for i, _ := range fruits {
		fmt.Println("Index ", i)
	}
	// or
	for i := range fruits { //the preferred way
		fmt.Println("Index ", i)
	}

	fruits2 := make([]string, 2)
	fruits2[0] = "apple"
	fruits2[1] = "manggo"
	fmt.Println(fruits2) // [apple manggo]
}
