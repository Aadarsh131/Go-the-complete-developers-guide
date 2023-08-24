package main

import "fmt"

type bot interface {
	getGreeting() string //any type that have a function named getGreeting() and is returning a string, is valid interface bot
}

type englishBot struct{} //in Go we have, implicit type interfaces(notice, we donot explicitly had specify which type is implementing which inteface, it is doing that with a just a function name, and that function internally will check its reciever to know if that type will be of our defined interface
type chineseBot struct{}

func main() {
	printGreeting(englishBot{})
	printGreeting(chineseBot{})
}

func (englishBot) getGreeting() string { //if we are not using the variable of the reciever type, then we can completely skip it
	return "Goodbye"
}

func (chineseBot) getGreeting() string {
	return "zao jian"
}

func printGreeting(b bot) { //we cannot overload the function declaration in Go, unlike other languages :( , but we can use the interface like this
	//this would mean, every agrument is passed if it is of type bot, in our case, englishbot{} and chineseBot{} satisfies the interface condition, hence both those types are valid argument
	//NOTE: interfaces are not generic programming
	fmt.Println(b.getGreeting())
}
