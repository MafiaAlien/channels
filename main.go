package main

import (
	"net/http"
	"os"
	// "io"
	"fmt"
)

func main(){
	links := []string {
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	for _, link := range links {
			go chekcLink(link)
		}
		
	
	}

func chekcLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		os.Exit(1)
		return 
	}

	fmt.Println(link, " is up!")

}