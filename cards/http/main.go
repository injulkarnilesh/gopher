package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://cat-fact.herokuapp.com/facts/")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	fmt.Println("Status", resp.Status)

	// data := make([]byte, 10000)
	// size, err := resp.Body.Read(data)
	// body := string(data)
	// fmt.Println(body)
	// fmt.Println("Body read size of ", size)

	//**********************

	// totalBody := 0
	// for {
	// 	data := make([]byte, 100)
	// 	size, err := resp.Body.Read(data)
	// 	body := string(data)
	// 	fmt.Println(body)
	// 	fmt.Println("Body read size of ", size)
	// 	totalBody += size

	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		fmt.Println("Error reading body", err)
	// 		os.Exit(1)
	// 	}
	// }
	// fmt.Println("Total Body read size of ", totalBody)

	//******************

	size, _ := io.Copy(os.Stdout, resp.Body)
	fmt.Println("Read size", size)

}
