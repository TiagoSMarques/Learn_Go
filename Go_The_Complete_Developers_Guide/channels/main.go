package main

import (
	"fmt"
	"net/http"
	"time"
	// "time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	//there is some delay when printing the result because for each request we are wating for the response to come back and we do not progress on through the function until the current request as been recieved
	// for _, link := range links {
	// 	checkDomain(link)
	// }

	//we can take a paralel aproach to this program
	// to do this we can make use of go routines and channels
	for _, link := range links {
		go checkDomain(link, c)
	}

	//	for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//	}

	// for {
	// 	go checkDomain(<-c, c)
	// }

	// this for loop is functionally the same as the one above
	// using range with a channel means wait for the channel to return some value and assing it to l
	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkDomain(link, c)
		}(l)
	}
}

// checks if the domain responds to requests
func checkDomain(link string, c chan string) {
	_, err := http.Get(link) //this is the blocking call

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
