package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please pass time argument")
		os.Exit(1)
	}
	timeStr := args[0]
	timeSec, err := strconv.Atoi(timeStr)
	if err != nil {
		fmt.Println("Can't convert to int:", timeStr)
		os.Exit(1)
	}

	signal := make(chan string)

	go func(sec int, ch chan string) {
		time.Sleep(time.Second * time.Duration(sec))
		ch <- "Done method call"
	}(5, signal)

	select {
	case message := <-signal:
		fmt.Println("Func call completed", message)
	case timedOut := <-time.After(time.Second * time.Duration(timeSec)):
		fmt.Print("Func call timed out at", timedOut)
	}

}
