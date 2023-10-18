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
	printGreetingVariadic(englishBot{}, chineseBot{})
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

// using a variadic function with ...bot or ...interface{}
func printGreetingVariadic(b ...interface{}) { // either use ...bot type, OR,  use ...interface{} but we have to use `type assertion` in this case
	for _, item := range b {
		fmt.Println(item.(bot).getGreeting()) //type casting would NOT work here, because we cannot convert a `interface{}` type to any other type, because `interface{}` is like a base abstract type in OOPS. BUT, we can convert int,float,bot etc types to `interface{}` type because these types have implemented all methods that `interface{}` has.
		//type assertion checks the type at the runtime
	}
}

//`any` or `interface{}` is not the `any` in typescript.
//`any`(an alias to `interface{}`) in Go is a type which allows you to assign any value to the variable, but you will never be able to call any method on the variable because the complier never knows
//`type assertion` is a thing in Go, which we can use here.
/*
Officially: `Type assertions` are used to check that a variable is of some type and return the underlying interface value. Type assertions work only for interfaces. For example, in the following code: 'var x interface{} = 42 t := x.(int)', 'x' has the 'interface{}' type with the underlying int value ('42'), 'int' is the concrete type that we want to check. If we print 't', the output will be '42'. Changing of the concrete type to 'string' ('t := x.(string)') will cause a runtime panic.
*/

////////////////I don't know what it means, but still keeping it for reference (random help from go discord server)///////////////
//So when you saw type cast, you know the target type won't have more details than the variable, type cast will always success (unlike Java).
//Type assert is actually more like the type cast in Java, you know the target type has more details than the variable, so you cannot ensure it will always success
//a specific case is unsafe.Pointer, you can cast *int64 to *string with the unsafe.Pointer in the middle, but that doesn't mean int64 can be cast to string. It means all the pointers have same underlying data structure
