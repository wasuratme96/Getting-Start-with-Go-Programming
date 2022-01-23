package main

import "fmt"

// Interface Type
type bot interface {
	getGreeting() string
}

// Concrete Type
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Assume this function is customer logic
	// for generating greeting in English
	return "Hi There !"
}

func (spanishBot) getGreeting() string {
	// for generating greeting in Spanish
	return "Hola !"
}
