package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign Values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Decalre and Assign
	foodArr := [2]string{"Rice", "Noodles"}

	// Slices
	fruitSlice := []string{"Grape", "Guava", "Apple"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(foodArr)
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
