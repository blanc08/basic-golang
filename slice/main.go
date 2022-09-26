package main

import "fmt"

func main() {
	// * spread
	var computers []string

	computers = append(computers, "Asus")
	computers = append(computers, "HP")
	computers = append(computers, "Lenovo")

	// computers := []string{"Asus"}

	fmt.Println(computers)

}
