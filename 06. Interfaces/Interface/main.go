package main

import (
	"fmt"
	"reflect"
)

type bot interface {
	getGreeting() string //any type that have a function named getGreeting() and is returning a string, is valid interface bot
}

type englishBot struct{} //in Go we have, implicit type interfaces(notice, we donot explicitly had to specify which type is implementing which interface, it is doing that with a just a function name
type chineseBot struct{}

func main() {
	printGreeting(englishBot{})
	printGreeting(chineseBot{})
	printGreeting_variadic(englishBot{}, chineseBot{})
}

func (englishBot) getGreeting() string { //if we are not using the variable of the reciever type, then we can completely skip it
	return "Goodbye"
}

func (chineseBot) getGreeting() string {
	return "zao jian"
}

func printGreeting(b bot) { //we cannot overload the function declaration in Go, unlike other languages :( , but we can use the interface like this :)
	//this would mean, every agrument is passed if it is of type bot, in our case, englishbot{} and chineseBot{} satisfies the interface condition, hence both those types are valid argument
	//NOTE: interfaces are not generic programming
	fmt.Println(b.getGreeting(), "(", reflect.TypeOf(b), ")")
	//Advaced: Or instead use a variadic function with param type ...bot (which is nothing but a slice type)
}

// using a variadic function with ...bot
func printGreeting_variadic(b ...bot) {
	for _, item := range b {
		fmt.Println(item.getGreeting(), "(", reflect.TypeOf(item), ")")
	}
}
