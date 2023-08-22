package main

func main() {
	//fmt.Println("Cards")
	//var card string = "Spades(ACE)"
	//card := "Spades(ACE)" //:= is only for initialization (first time declaration), cannot be used for assigning a new value after the initialization
	//Go can help us with type inferencing and can auto deduce the type
	//card = "Diamond(Five)" // = for assignment
	//card := newCard()
	//we have slice(a dymaic sized array) and array(fixed size) in Go
	//cards := deck{newCard(), "Diamnod(ACE)"} //slice

	//append into slice
	//cards = append(cards, "lol") //not modifying the original slice, a new copy of cards has been made

	//cards.print() //In OO programming terms - it is only possible beacuse we "cards" is of type deck, and we have a function in type "deck" as "print()"
	//In Go programming terms- It was possible because we had a reciever(of type deck) on the function print()

	newDeck().print()
	//OR
	//cards := newDeck()
	//cards.print()
}

func newCard() string {
	return "Diamond(Five)"
}
