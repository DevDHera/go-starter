package main

import "fmt"

func main() {
	ids := []int{33, 44, 55, 66, 2, 1, 67}

	// Loop through Ids
	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum =  %d\n", sum)

	// Range with Maps
	studentIds := map[int]string{1: "Devin", 2: "Don"}

	for m, n := range studentIds {
		fmt.Printf("%d: %s\n", m, n)
	}

	for m := range studentIds {
		fmt.Println("Key = ", m)
	}
}
