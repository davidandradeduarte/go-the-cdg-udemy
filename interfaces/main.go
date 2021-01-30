package main

import "fmt"

type bot interface {
	getGreeting() string
	xpto()
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	fmt.Println(eb)
	fmt.Println(sb)
	// printGreeting(eb)
	// printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// VERY custome logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// VERY custome logic for generating a spanish greeting
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
