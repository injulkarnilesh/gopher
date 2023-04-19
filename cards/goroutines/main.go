package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://amazon.com",
		"http://twader.com",
		"http://stackoverflow.com",
		"http://fast.com",
	}
	ch := make(chan siteStatus)
	for _, site := range sites {
		go ping(site, ch)
	}

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	for st := range ch {
		fmt.Println(st)
	}
}

func ping(site string, ch chan siteStatus) {
	_, err := http.Get(site)

	if err != nil {
		ch <- siteStatus{site: site, status: "down"}
		return
	}

	ch <- siteStatus{site: site, status: "up"}
}

type siteStatus struct {
	site   string
	status string
}
