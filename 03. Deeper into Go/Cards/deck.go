package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// create a new type of 'deck' which is a slice of string

type deck []string // in oo programming terms, 'deck' here is a class that borrows its functionalities from the []string(a slice of strings)

func newDeck() deck { //a function without a reciever can be seen as a function that doesn't attach to any instance, its a normal function
	// cards := deck{}
	//OR
	var cards []string
	cardsuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardnum := []string{"Ace", "two", "three", "four"}

	for _, suit := range cardsuits {
		for _, num := range cardnum {
			cards = append(cards, (num + "(" + suit + ")"))
			//fmt.Println(cards)
		}
	}
	return cards //here cards of type []string is implicitly being type casted to type deck
	//OR
	// return deck(cards)
}

// in Go, it is called a reciever (a convention to use a single character that matches the type name, eg 'd' in our case)
func (d deck) print() { //in oo programming terms, this function is specific to the deck instance 'd' (where d can be any variable name[instance/reference] that uses deck class),(it is like "this" keyword), and not to the base class []string (means any variable/instance of type deck will now have the access to print() function )
	for i, card := range d { // := is used because internally Go will delete the previous variables while iterating (maybe an example of garbage collection, or for performance reasons), so each time we have to initialize a new varible and hence := is used
		fmt.Print(i, card, " ")
	}
	fmt.Println()
}

func deal(d deck, handsize int) (deck, deck) { //returning multiple values of both deck type
	return d[:handsize], d[handsize:]
}

// converts a type deck to string
func (d deck) toString() string {
	//var result string
	//for _, card := range d {
	//	result += (card + ",")
	//}
	//return result

	//OR

	//using the in-built strings library method for more ease
	//strSlice := []string(d) //we can convert back a deck type to []string(although internally deck is of type []string)

	//OR simply
	return strings.Join(d, ",") //as type of d->deck->[]string, and type []string is what we actually need here
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading file: ", err)
		os.Exit(1)
	}
	stringFile := string(byteSlice)
	//stringSlice := strings.Split(stringFile, ",")
	//return deck(stringSlice) //converted []string to deck(which is internally the same datatype, hence we can use the below statement too, and Go compiler will figure out and convert a []string to deck internally

	//OR
	return strings.Split(stringFile, ",") //although Split() method returns []string, it internally must be converting to deck type, that is the same reason we are able to use print() method(our custom method) on the returned value from this function (to be specific, this line- readFile("deal1.txt").print() from the main function)
}

func (d deck) shuffle() {
	for i := range d {
		newPos := rand.Intn(len(d))
		d[i], d[newPos] = d[newPos], d[i] //swap
	}
}
