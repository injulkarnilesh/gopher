package main

import "fmt"

func mainOddEven() {
	s := []int{}
	for i := 0; i <= 10; i++ {
		s = append(s, i)
	}
	for _, i := range s {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}

func mainCards() {
	//printState()
	card := newName()
	fmt.Println(card)

	//cards := [] string { "Five of Diamond", newName() }
	// cards := deck { "Five of Diamond", newName() }
	// cards = append(cards, "Ten of Hearts")
	cards := NewDeck()
	fmt.Println("Original Cards")
	cards.print()

	hand, restOfCards := deal(cards, 5)
	fmt.Println("Hand Cards")
	hand.print()

	fmt.Println("Rest of Cards")
	restOfCards.print()

	fmt.Println("toString rest of Cards")
	fmt.Println(restOfCards.toString())

	const FILE_NAME = "my-cards"
	restOfCards.saveToFile(FILE_NAME)

	savedCards := newDeckFromFile(FILE_NAME)
	fmt.Println("Cards from file")

	savedCards.shuffle()
	savedCards.print()

	// for i, theCard := range cards {
	// 	fmt.Println(i, theCard)
	// }

}

func main() {
	nilesh := person{
		firstName: "Nielsh",
		lastName:  "Injulkar",
	}
	fmt.Println(nilesh)
	nilesh.updateFirstName("nils")
	nilesh.updateLastName("inju")

	fmt.Println(nilesh)

	(&nilesh).updateLastName("injukar")
	fmt.Println(nilesh)
}

func newName() string {
	return "Ace of Spades"
}
