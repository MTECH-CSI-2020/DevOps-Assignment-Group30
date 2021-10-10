package main

func main() {
	cards := newDeck()
	//cards := newDeckFromFile("Test.txt")
	cards.shuffle()
	cards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("Test.txt")
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()
}
