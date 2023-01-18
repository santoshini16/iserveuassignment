package main

import "fmt"

type bot interface {
	getgreeting() string
}
type englishbot struct{}
type spanishbot struct{}

func main() {
	eb := englishbot{}
	sb := spanishbot{}

	printgreeting(eb)
	printgreeting(sb)
}

func printgreeting(b bot) {
	fmt.Println(b.getgreeting())
}

func (englishbot) getgreeting() string {
	//VERY custom logic for generating an english greeting
	return "Hi There!"
}
func (spanishbot) getgreeting() string {
	return "Hola!"
}
