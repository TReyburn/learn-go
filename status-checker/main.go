package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://totallyfake.abc",
		"http://alteryx.com",
		"http://reddit.com",
		"http://colorado.edu",
		"http://www.solveigdelabroye.me",
	}

	for _, u := range urls {
		checkUrl(u)
	}
}

func checkUrl(url string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println("Status:", url, "might be down.")
		return
	}

	fmt.Println("Status:", url, "is up.")
}