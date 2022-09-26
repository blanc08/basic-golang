package main

import "fmt"

func main() {
	// var myMap map[string]int
	// myMap = map[string]int{}

	// * shorter syntax
	// myMap := make(map[string]int)

	// myMap["Ruby"] = 1

	// * direct
	// myMap := map[string]float64{"a": 1.2, "b": 2.2}

	// * for loop
	// for _, v := range myMap {
	// fmt.Println(v)
	// }

	// * check existence
	// value, isAvailable := myMap["c"]
	// fmt.Print(value, isAvailable)
	// fmt.Println(myMap)

	// * slice map
	students := []map[string]string{
		{"name": "Agung", "score": "A"},
		{"name": "Ilham", "score": "B"},
		{"name": "Aris", "score": "C"},
	}

	for _, values := range students {
		fmt.Println(values["name"])
	}

	fmt.Println(students)

}
