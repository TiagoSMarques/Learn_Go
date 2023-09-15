package main

import "fmt"

type bot interface {
	//this means that if any type inside this file as the function "getGreeting" this type is also of type bot
	getGreeting() string
}

type englishBot struct{}
type portugueseBot struct{}

func main() {
	eb := englishBot{}
	pb := portugueseBot{}

	printGreeting(eb)
	printGreeting(pb)
}

// this function may be called by every element of type bot
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic to generate a english greeting
	return "Hi there!"
}

func (pb portugueseBot) getGreeting() string {
	// very custom logic to generate a portuguese greeting
	return "Tá tudo?"
}
