package main

import (
	"fmt"
	"os"
)

func main() {
	argv := os.Args
	for _, a:= range argv {
		fmt.Println(a)
	}
}
