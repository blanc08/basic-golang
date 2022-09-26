package main

import "fmt"

func main() {

	// catching multiple return function
	luas, keliling := calculate(2, 4)

	fmt.Println(luas, keliling)

	return
}

// * multiple return function
func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}
