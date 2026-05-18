package main

import "fmt"

func main() {
	fruits := []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	fruitsA := []string{"apple", "grape"}     // slice
	fruitsB := [2]string{"banana", "melon"}   // array
	fruitsC := [...]string{"papaya", "grape"} // array

	fmt.Printf("The data types of variable fruitsA is :%T\n", fruitsA)
	fmt.Printf("The data types of variable fruitsB is :%T\n", fruitsB)
	fmt.Printf("The data types of variable fruitsC is :%T\n", fruitsC)

	fmt.Println(fruits[0:2]) // ["apple", "grape"]
	fmt.Println(fruits[0:4]) // [apple, grape, banana, melon]
	fmt.Println(fruits[0:0]) // []
	fmt.Println(fruits[4:4]) // []
	// fmt.Println(fruits[4:0])	//this line will throws an error since [a:b] a should be smaller than b
	fmt.Println(fruits[:])  // [apple, grape, banana, melon]
	fmt.Println(fruits[2:]) // [banana, melon]
	fmt.Println(fruits[:2]) // [apple, grape]

	fmt.Println() //slice is a data reference type
	aFruits := fruits[0:3]
	bFruits := fruits[1:4]
	aaFruits := aFruits[1:2]
	baFruits := bFruits[0:1]
	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	baFruits[0] = "pineapple" // "grape" is changed to "pineapple"
	fmt.Println(fruits)       // [apple pineapple banana melon]
	fmt.Println(aFruits)      // [apple pineapple banana]
	fmt.Println(bFruits)      // [pineapple banana melon]
	fmt.Println(aaFruits)     // [pineapple]
	fmt.Println(baFruits)     // [pineapple]

	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	cFruits := append(fruits, "papaya")
	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]
	bFruits = fruits[0:2]
	fmt.Println(cap(bFruits)) // 3
	fmt.Println(len(bFruits)) // 2
	fmt.Println(fruits)       // ["apple", "grape", "banana"]
	fmt.Println(bFruits)      // ["apple", "grape"]
	// cFruits = append(bFruits, "papaya")
	fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruits) // ["apple", "grape"]
	fmt.Println(cFruits) // ["apple", "grape", "papaya"]

	dst := make([]string, 3) //dst has smaller number of elements than src
	src := []string{"watermelon", "pineapple", "apple", "orange"}
	n := copy(dst, src)
	fmt.Println(dst) // watermelon pineapple apple
	fmt.Println(src) // watermelon pineapple apple orange
	fmt.Println(n)   // 3 (the return of the copy function is the minimum len() of src and dst)

	dst = []string{"potato", "potato", "potato"} //dst has bigger number of elements than src
	src = []string{"watermelon", "pineapple"}
	n = copy(dst, src)
	fmt.Println(dst) // watermelon pineapple potato
	fmt.Println(src) // watermelon pineapple
	fmt.Println(n)   // 2

	fruits = []string{"apple", "grape", "banana"}
	aFruits = fruits[0:2]
	bFruits = fruits[0:2:2]
	fmt.Println(fruits)       // ["apple", "grape", "banana"]
	fmt.Println(len(fruits))  // len: 3
	fmt.Println(cap(fruits))  // cap: 3
	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3
	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
}
