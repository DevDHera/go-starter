package main

import "fmt"

func main() {
	x := 14
	y := 6

	// If else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// If else if
	color := "Blue"

	if color == "Red" {
		fmt.Println("Color is Red")
	} else if color == "Blue" {
		fmt.Println("Color is Blue")
	} else {
		fmt.Println("Color is not equal to Red or Blue")
	}

	// Switch
	switch color {
	case "Red":
		fmt.Println("Color is Red")
	case "Blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is not equal to Red or Blue")
	}
}
