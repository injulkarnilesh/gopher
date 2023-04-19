package main

import (
	"fmt"
	"strconv"
	"time"
)

func mainSimple() {
	ch := make(chan string)

	ch <- "Nilesh"
	name := <-ch

	fmt.Println("Received", name)
}

func main() {
	count := 5
	buffch := make(chan string)
	for i := 0; i < count; i++ {
		go routine(i, buffch)
	}
	time.Sleep(time.Second)

	fmt.Println("Reading ", 0)
	fmt.Println(<-buffch)
	fmt.Println("Read ", 0)

	fmt.Println("Reading ", 1)
	fmt.Println(<-buffch)
	fmt.Println("Read ", 1)

	time.Sleep(time.Second)
}

func routine(num int, ch chan string) {
	name := strconv.Itoa(num)
	fmt.Println("Called ", name)
	fmt.Println("Sending ", name)
	ch <- "Done " + name
	fmt.Println("Sent ", name)
}
