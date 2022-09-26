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
	// students := []map[string]string{
	// 	{"name": "Agung", "score": "A"},
	// 	{"name": "Ilham", "score": "B"},
	// 	{"name": "Aris", "score": "C"},
	// }

	// for _, values := range students {
	// 	fmt.Println(values["name"])
	// }

	// fmt.Println(students)

	// * Quiz
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// * rata rata
	rate := getRate(scores[:])
	fmt.Println(rate)

	// * good scores
	goodScores := getGoodScores(scores[:])
	fmt.Println(goodScores)
}

func getGoodScores(scores []int) []int {
	var tempScores []int

	for _, value := range scores {
		if value >= 90 {
			tempScores = append(tempScores, value)
		}
	}

	return tempScores
}

func getRate(scores []int) float64 {
	var total float64
	for _, value := range scores {
		total += float64(value)
	}

	return total / float64(len(scores))
}
