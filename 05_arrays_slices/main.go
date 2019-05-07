package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assing values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assing
	nameArray := [2]string{"Felipe", "Campara"}
	ageSlice := []int{26, 30}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(nameArray)
	fmt.Println(ageSlice)
	fmt.Println(len(ageSlice))
	fmt.Println(nameArray[0:1])
}
