package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number", i)
	}

	i := 0
	for i < 5 {
		fmt.Println("Number", i)
		i++
	}

	i = 0
	for {
		fmt.Println("Number", i)
		i++
		if i == 5 {
			break
		}
	}

	xs := "123"
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}

	ys := [5]int{10, 20, 30, 40, 50}
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	zs := ys[0:2]
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	kvs := map[byte]int{'a': 0, 'b': 1, 'c': 2}
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}

	for range kvs {
		fmt.Println("Done")
	}

	for i := range 5 {
		fmt.Println(i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Value", i)
	}

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
				// continue outerLoop
			}
			fmt.Print("Matrix [", i, "][", j, "]", "\n")
		}
	}
}
