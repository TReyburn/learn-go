package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff000",
		"green": "#4b781",
	}
	delete(colors, "red")
	fmt.Println(colors)

	var otherColors map[string]string
	fmt.Println(otherColors)

	anotherColors := make(map[string]string)
	anotherColors["black"] = "#0000"
	fmt.Println(anotherColors)
}
