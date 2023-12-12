package main

import "fmt"

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

	//cards.print() //In OO programming terms - it is only possible beacuse "cards" is of type deck, and we have a function in type "deck" as "print()"
	//In Go programming terms- It was possible because we had a reciever(of type deck) on the function print(), so all the deck type variables now have the access to the print() function

	newDeck().print()
	//OR
	//cards := newDeck()
	//cards.print()

	fmt.Println("---------------------Make a Deal-------------------------")
	a, b := deal(newDeck(), 3)
	a.print()
	b.print()

	//conversion string to []byte
	//fmt.Println([]byte("helloww")) // []byte() will convert a string to []byte

	//to convert a []string to []byte, first convert []string to string, then we can easily convert it to []byte from string
	//fmt.Println(a.toString())
	fmt.Println([]byte(a.toString())) //incl. commas

	err := a.saveToFile("deal1.txt")
	if err != nil {
		fmt.Println("couldn't save, because of error: ", err)
	}
	err = b.saveToFile("deal2.txt")
	if err != nil {
		fmt.Println("couldn't save, because of error: ", err)
	}

	fmt.Println("=======================READ File=============================")
	readFile("deal1.txt").print()

	fmt.Println("-----------------------------------------------------------")
	b.shuffle()
	readFile("deal2.txt").print()

}

//func newCard() string {
//	return "Diamond(Five)"
//}
