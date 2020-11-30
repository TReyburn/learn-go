package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(resp)

	//bs := make([]byte, 99999)
	//_, err = resp.Body.Read(bs)
	//fmt.Println(string(bs))

	lw := logWriter{}

	_, _ = io.Copy(lw, resp.Body)

}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
