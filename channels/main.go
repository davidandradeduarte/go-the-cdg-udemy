package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkURL(url, c)
	}

	for l := range c {
		go func(url string) {
			time.Sleep(time.Second * 5)
			checkURL(url, c)
		}(l)
	}
}

func checkURL(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}
