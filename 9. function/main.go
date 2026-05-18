package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type tryType func(int) bool

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	names := []string{"John", "Wick"}
	printMessage("Hello", names)

	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
	divideNumber(randomWithRange(0, 10), randomWithRange(0, 10))

	var cDia float64 = 7
	area, circumference := calculate(cDia)
	fmt.Printf("The calculations of %.3f cm circle are:\nArea\t\t: %.3f cm\nCircumference\t: %.3f cm\n", cDia, area, circumference)
	area, circumference = calculate2(cDia)
	fmt.Printf("The calculations of %.3f cm circle are:\nArea\t\t: %.3f cm\nCircumference\t: %.3f cm\n", cDia, area, circumference)

	avg := calculate3(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	msg := fmt.Sprintf("Average value\t: %.2f", avg)
	fmt.Println(msg)

	numbers2 := []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	avg2 := calculate3(numbers2...)
	msg2 := fmt.Sprintf("Average value\t: %.2f", avg2)
	fmt.Println(msg2)

	yourHobbies("wick", "sleeping", "eating")

	getMinMax := func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	min, max := getMinMax(numbers2)
	fmt.Printf("%v numbers have a minimum value of %v and a maximum value of %v\n", numbers2, min, max)

	newNumbers := func(min int) []int {
		var r []int
		for _, e := range numbers2 {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)
	fmt.Println("original number :", numbers2)
	fmt.Println("filtered number :", newNumbers)

	setMax := 4
	numsMaxNum, maxNumNums := findMax(numbers2, setMax)
	fmt.Printf("%v have %v (%v) that are equal or below %v", numbers2, numsMaxNum, maxNumNums(), setMax)

	data := []string{"wick", "jason", "ethan"}
	dataContainsO := filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	dataLenght5 := filter(data, func(each string) bool {
		return len(each) == 5
	})
	fmt.Println("data \t\t:", data)
	fmt.Println("filter character of\"o\"\t:", dataContainsO)
	fmt.Println("filter len of\"5\"\t:", dataLenght5)

	maxNumFunc := iTry(numbers2, func(numbers int) bool {
		return numbers <= 4
	})

	fmt.Println(maxNumFunc)

	dataTry := []string{"Los", "Pollos", "Hermanos", "Gus", "Fring"}
	containO := anotherTry(dataTry, func(word string) bool {
		return strings.Contains(word, "u")
	})
	fmt.Println(containO)

}

func printMessage(message string, arr []string) {
	// nameString := strings.Join(arr, " ")
	// fmt.Printf("Message : \"%s %s\"", message, nameString)
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	value := randomizer.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}
	res := m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

func calculate(d float64) (float64, float64) {
	area := math.Pi * math.Pow(d/2, 2)
	circumference := math.Pi * d
	return area, circumference
}

func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return
}

func calculate3(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}
	avg := float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, ", ")
	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func iTry(data []int, callback tryType) []int {
	var result []int
	for _, each := range data {
		if funct := callback(each); funct {
			result = append(result, each)
		}
	}
	return result
} // Using alias in Closure types

func anotherTry(data []string, callback func(string) bool) []string {
	var res []string
	for _, each := range data {
		if condition := callback(each); condition {
			res = append(res, each)
		}
	}
	return res
}
