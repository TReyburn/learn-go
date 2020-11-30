package main

import (
	"fmt"
	"io"
	"os"
)

type fileContent struct {
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Need to include a file to be read")
		os.Exit(1)
	}

	fn := os.Args[1]
	fc := fileContent{}
	fh, err := os.Open(fn)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = io.Copy(fc, fh)
}

func (fc fileContent) Write(b []byte) (n int, err error) {
	fmt.Println(string(b))
	return len(b), nil
}
