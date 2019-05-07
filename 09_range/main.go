package main

import "fmt"

func main() {
	ids := []int{37, 36, 23, 67, 32, 1, 6}

	// Loop throug ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{
		"Bob":    "bob@gmail.com",
		"Felipe": "felipe@gmail.com",
	}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}

	for key := range emails {
		fmt.Println("Name: " + key)
	}
}
