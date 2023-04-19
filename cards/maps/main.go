package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "#ff0000",
		"green":  "#78aa00",
		"yellow": "#3fabcaa",
	}
	fmt.Println(colors)
	printMap(colors)

	var myColors map[string]int = make(map[string]int)
	myColors["white"] = 0
	myColors["black"] = 10
	fmt.Println(myColors)

	delete(myColors, "white")
	fmt.Println(myColors)

}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("KV", k, v)
	}
}
