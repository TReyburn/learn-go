package main

import "fmt"

func main() {
	// Create our int slice from 0 to 10
	var s []int
	for i := 0; i <= 10; i++ {
		s = append(s, i)
	}

	// Using modulus to check if number is even or odd for each int in our slice
	for _, n := range s {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}
}
