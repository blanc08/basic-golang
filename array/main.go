package main

import "fmt"

func main() {
	// * normal initalization
	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "Javascript"
	// languages[3] = "C"
	// languages[4] = "Python"

	// * try to allocating on out of bound index
	// languages[5] = "Error"

	// * direct way
	// languages := [5]string{"Go", "Ruby", "Javascript", "C", "Python"}

	// * vertical initalization
	// languages := [5]string{
	// 	"Go",
	// 	"Ruby",
	// 	"Javascript",
	// 	"C",
	// 	"Python",
	// }

	// * spread
	languages := [...]string{
		"Go",
		"Ruby",
		"Javascript",
		"C",
		"Python",
		"Kotlin",
	}

	// * looping
	for _, language := range languages {
		fmt.Println(language)
	}

	fmt.Println(languages)
	length := len(languages)
	fmt.Println(length)

}
