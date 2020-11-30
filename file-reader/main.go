package main

import (
	"fmt"
	"io"
	"os"
)

type fileContent struct {
}

func main() {
	fn := os.Args[1]
	fc := fileContent{}
	fh, err := os.Open(fn); if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = io.Copy(fc, fh)
}

func (fc fileContent) Write(b []byte) (n int, err error) {
	fmt.Println(string(b))
	return len(b), nil
}
