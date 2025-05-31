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

	var pointGrade = 8840.0
	if percent := pointGrade / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var pointWord = 6
	switch pointWord {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	switch pointWord {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	pointWord = 3
	switch pointWord {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	switch {
	case pointWord == 8:
		fmt.Println("perfect")
	case (pointWord < 8) && (pointWord > 2):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	pointWord = 6
	switch {
	case pointWord == 8:
		fmt.Println("perfect")
	case (pointWord < 8) && (pointWord > 3):
		fmt.Println("awesome")
		fallthrough
	case pointWord < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	pointWord = 9
	if pointWord > 7 {
		switch pointWord {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if pointWord == 5 {
			fmt.Println("not bad")
		} else if pointWord == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if pointWord == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
