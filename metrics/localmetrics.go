package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type MetricsStore struct {
	d                 time.Duration
	store             map[int64]map[string]map[string]float64
	startTime         int64
	lastProcessedTime int64
	ticker            *time.Ticker
}

func NewMetricsStore(d time.Duration) *MetricsStore {
	emptyMap := make(map[int64]map[string]map[string]float64)

	met := MetricsStore{
		d:     d,
		store: emptyMap,
	}
	go startConsuming(&met)
	return &met
}

func startConsuming(s *MetricsStore) *time.Ticker {
	ticker := time.NewTicker(s.d)
	s.ticker = ticker
	for {
		select {
		case <-ticker.C:
			consume(s)
		}
	}
}

func consumeRest(s *MetricsStore) {
	if s.ticker != nil {
		s.ticker.Stop()
	}
	for !s.IsEmpty() {
		consume(s)
	}
}

func consume(s *MetricsStore) {
	fmt.Println("Consuming data")
	if s.startTime != 0 {
		var toProcess int64
		if s.lastProcessedTime == 0 {
			toProcess = s.startTime
		} else {
			toProcess = s.lastProcessedTime + s.d.Milliseconds()
		}

		currentWindow := time.Now().Truncate(s.d).UnixMilli()
		canProcess := currentWindow-toProcess > s.d.Milliseconds()

		if canProcess {
			if data, ok := s.store[toProcess]; ok {
				printData(toProcess, data)
				s.lastProcessedTime = toProcess
				delete(s.store, toProcess)
			} else {
				fmt.Println("No data for time ", time.Unix(0, toProcess*int64(time.Millisecond)))
			}
		} else {
			fmt.Println("Current window yet to be filled")
		}
	} else {
		fmt.Println("No data to consume")
	}
}

func (s *MetricsStore) IsEmpty() bool {
	return len(s.store) == 0
}

func (s *MetricsStore) Print() {
	for t, values := range s.store {
		fmt.Println("From time ", time.Unix(0, t*int64(time.Millisecond)))
		for name, labels := range values {
			fmt.Println("  Name", name)
			for label, value := range labels {
				fmt.Println("    ", label, "=", value)
			}
		}
	}
}

func printData(t int64, values map[string]map[string]float64) {
	fmt.Println("From time ", time.Unix(0, t*int64(time.Millisecond)))
	for name, labels := range values {
		fmt.Println("  Name", name)
		for label, value := range labels {
			fmt.Println("    ", label, "=", value)
		}
	}
}

func (s *MetricsStore) AddCounter(name string, count float64, labels []string, labelValues []string) {

	if len(labels) != len(labelValues) {
		panic("Label and label values count mismatch")
	}
	valuedLabels := []string{}
	for i, label := range labels {
		valuedLabels = append(valuedLabels, label+"="+labelValues[i])
	}
	sort.Strings(valuedLabels)
	labelKey := strings.Join(valuedLabels, "#")

	millis := time.Now().Truncate(s.d).UnixMilli()
	if s.startTime == 0 {
		s.startTime = millis
	}

	byTimeValue, ok := s.store[millis]

	if !ok {
		byTimeValue = make(map[string]map[string]float64)
		s.store[millis] = byTimeValue
	}

	byNameValue, ok := byTimeValue[name]

	if !ok {
		byNameValue = make(map[string]float64)
		byTimeValue[name] = byNameValue
	}

	byLabelValue, ok := byNameValue[labelKey]
	if !ok {
		byNameValue[labelKey] = 1
	} else {
		byNameValue[labelKey] = byLabelValue + count
	}

}
