package main

import (
	"fmt"
	"sync"
	"time"
)

func conc() {

	// this code will never count fish because count sheep never ends
	count("sheep")
	count("fish")
	// we can make the funct count run in another thread by adding go
	go count("sheep")
	count("fish")
	// if we add go to both funtions the progam wont execte anything because in go when the program hits the end of the main funt it ends

	//to wait for a funtion to finish we can use a wait group
	var wg sync.WaitGroup
	wg.Add(1) //this signifies that we have 1 go rotine to wait for

	//wrap the count funtion in an anonymous funtion that decrements the wg counter when it ends

	go func() {
		count("sheep")
		wg.Done()
	}()
	wg.Wait() //this will wait until the wg counter is 0
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

//to make this usefull we can create channel as a way of comunication beteween concurrent funtions

func conc2() {
	c := make(chan string)
	go count2("sheep", c)

	for {
		msg, open := <-c //this is a blocking operation meaning we have to wait for something to recieve
		if !open {
			break
		}

		fmt.Println(msg)
	}

	//a nicer way of doing the above:
	for msg := range c {
		fmt.Println(msg)
	}

}

func count2(thing string, c chan string) {

	for i := 1; i <= 5; i++ {
		c <- thing //passing our input to the channel
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // if you are the sender and there is nothing more to send we should close the channel, else we are going to run into a runetime error because the reciver is waiting for a message that will never come
}

func conc3() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 s"
			time.Sleep(time.Second * 2)
		}
	}()

	//since recieving a message is a blocking operation we have to wait forthe 2s message to arrive to procede with the first funtion meaning that the performance for the first is dependent on the second

	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	// to solve this we use select which recieves from the first channel that is ready

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// example of a common pattern called worker pulls were we have multiple concurrent workers pulling of a work queue
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// we can specify in the funtion defenition that jobs only RECIEVES and results can only SEND
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
