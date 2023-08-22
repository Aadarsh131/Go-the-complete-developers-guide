package main

import "fmt"

// create a new type of 'deck' which is a slice of string

type deck []string // in oo programming terms, 'deck' here is a class that borrows its functionalities from the []string(a slice of strings)

func newDeck() deck { //a function without a reciever can be seen as a function that doesn't attach to any instance, its a normal function
	cards := deck{}
	cardsuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardnum := []string{"Ace", "two", "three", "four"}

	for _, suit := range cardsuits {
		for _, num := range cardnum {
			cards = append(cards, (num + "(" + suit + ")"))
			//fmt.Println(cards)
		}
	}
	return cards
}

// in Go, it is called a reciever (a convention to use a single character that matches the type name, eg 'd' in our case)
func (d deck) print() { //in oo programming terms, this function is specific to the deck instance 'd' (where d can be any variable name[instance/reference] that uses deck class),(it is like "this" keyword), and not to the base class []string (means any variable/instance of type deck will now have the access to print() function )
	for i, card := range d { // := is used because internally Go will delete the previous variables while iterating (maybe an example of garbage collection), so each time we have to initialize a new varible and hence := is used
		fmt.Println(i, card)
	}
}
