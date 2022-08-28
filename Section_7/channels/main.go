package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://facebook.com",
		"https://www.ea.com/games",
		"https://www.google.co.in/",
		"https://www.google.com",
		"https://www.google.co.jp/",
		// "https://www.google.co.ca/",
		"https://www.google.com.mx",
		// "https://www.google.co.mx",
	}

	c := make(chan string)

	for _, l := range links {
		go checkStatus(l, c)
	}

	for l := range c {
		go func (link string) {
			time.Sleep(time.Second * 5)
			checkStatus(link, c)
		}(l)
	}
}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}