package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	m := NewMetricsStore(time.Second * 10)
	ticker := time.NewTicker(time.Second * 15)
	processes := []string{"Rule", "Parser", "Writer", "Ingester"}
	customers := []string{"Tesla", "Apple", "Google", "MicroSoft"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 0

	for {
		select {
		case <-ticker.C:
			fmt.Println("********")

			count++
			if count == 5 {
				//m.Print()
				consumeRest(m)
				os.Exit(0)
			}
		default:
			process := processes[r.Intn(len(processes))]
			customer := customers[r.Intn(len(customers))]

			m.AddCounter("events_dropped", 1, []string{"customer", "process"}, []string{customer, process})
			time.Sleep(time.Millisecond * 500)
		}
	}

}
