package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["January"] = 50
	chicken["February"] = 40
	fmt.Println("January", chicken["January"]) // January 50
	fmt.Println("May", chicken["May"])         // May 0

	// var data map[string]int
	// data["one"] = 1	//will throw an error since the default value is nil

	data := map[string]int{}
	data["one"] = 1
	fmt.Println(data["one"])

	chicken1 := map[string]int{"January": 50, "February": 40} // horizontal initialization

	chicken2 := map[string]int{ // vertical initialization
		"January":  50,
		"February": 40,
	}
	fmt.Println(chicken1["January"])
	fmt.Println(chicken2["January"])

	chicken3 := map[string]int{}
	chicken4 := make(map[string]int)
	chicken5 := *new(map[string]int)
	fmt.Println(chicken3["Los"])
	fmt.Println(chicken4["Polos"])
	fmt.Println(chicken5["Hermanos"])

	chicken6 := map[string]int{
		"January":  50,
		"February": 40,
		"March":    34,
		"April":    67,
	}
	for key, val := range chicken6 {
		fmt.Println(key, " \t:", val)
	}

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)
	delete(chicken, "January")
	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

	value, isExist := chicken["May"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// chickens7 := []map[string]string{
	// 	map[string]string{"name": "chicken blue", "gender": "male"},
	// 	map[string]string{"name": "chicken red", "gender": "male"},
	// 	map[string]string{"name": "chicken yellow", "gender": "female"},
	// }	//type redundancy
	chickens7 := []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickens7 {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	chickenData := []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	fmt.Println(chickenData[0]["name"])
	for _, data := range chickenData {
		fmt.Println(data)
	}
}
