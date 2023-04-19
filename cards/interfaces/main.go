package main

import (
	"fmt"
)

type bot interface {
	greeting() string
	//fullName(string, string) (string, error)
}

type robot interface {
	bot
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

	tr := triangle{
		base:   64,
		height: 12,
	}
	sq := square{
		sideLength: 10.1,
	}

	printArea(tr)
	printArea(sq)
}

func (englishBot) greeting() string {
	return "hi"
}

func (spanishBot) greeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.greeting())
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func (t triangle) area() float64 {
	return t.base * t.height * 0.5
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	area() float64
}

func printArea(s shape) {
	fmt.Println(s.area())
}
