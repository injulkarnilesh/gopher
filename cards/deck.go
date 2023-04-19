package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

type deck [] string

func NewDeck() deck {
	cards := deck{}
	cardSuits := []string { "Spades", "Diamonds", "Hearts", "Clubs" }
	cardValues := []string { "Ace", "One", "Two", "Three", "Four" }
	for _, value := range cardValues {
		for _, suit := range cardSuits {
			card := value + " of " + suit
			cards = append(cards, card) 
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, s int) (deck, deck) {
	return d[:s], d[s:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	ba, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	cardStrings := strings.Split(string(ba), ",")
	return deck(cardStrings)
}

func (d deck) shuffle() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)
	for i := range d {
		j := r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}