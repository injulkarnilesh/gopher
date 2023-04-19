package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	arguments := os.Args
	fmt.Println(arguments)
	if len(arguments) < 2 {
		fmt.Println("Please pass the fileName")
		os.Exit(1)
	}
	fileName := arguments[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to read file by name "+fileName, err)
		os.Exit(1)
	}
	copied, err := io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Failed to copy from file", err)
		os.Exit(1)

	}
	fmt.Printf("\n#########Copied %d bytes#########\n", copied)
}
