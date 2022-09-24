package main

import (
	"net/http"
	// "os"
	// "io"
	"fmt"
	"time"
)

func main(){
	links := []string {
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string) // create a channel transferring  string 

	for _, link := range links {
			go checkLink(link, c)
		}

	for l := range c {
		go func (link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

	}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return 
	}
	fmt.Println(link, " is up!")
	c <- link
}