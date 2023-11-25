package main

import "fmt"

type contact struct {
	email        string
	currectPlace string
}
type person struct {
	firstname string
	lastname  string
	age       int
	//contactInfo contact
	//OR
	contact //shorthand for: contact contact(a type contact, of variable name contact)
}

type SLICE []string

func main() {
	//aadarsh := person{"aadarsh", "shah", 23} //order must be same
	//aadarsh := person{firstname: "aadarsh", lastname: "shah", age: 23}

	var aadarsh person //assignes a zero-value by go compiler (string-> "", int-> 0, float->0, bool->false) corresponding zero-values for some common datatypes
	//aadarsh.firstname = "atul"
	//aadarsh.lastname = "shah"
	//aadarsh.age = 23
	//aadarsh.contactInfo = contact{email: "aadarshkumar131@gmail.com", currectPlace: "bagdogra"}
	//fmt.Printf("%+v", aadarsh) //prints the field name with its value

	//aadarsh := person{firstname: "aadarsh", lastname: "shah", contact: contact{email: "aadarshkumar131@gmail.com", currectPlace: "Bagdogra"}}
	//aadarsh.print()
	aadarsh.updateFirstName("atul")
	//aadarsh.print()

	temp := SLICE{"z", "b", "c"}
	temp.updateMe("a")
	fmt.Println(temp)
}

func (p *person) updateFirstName(newFirstName string) { //using pointer as a reciever, meaning the value will be pass by ref
	//NOTE: (p person) would make it pass by value(copy would be made)

	fmt.Println(&p)
	(*p).firstname = newFirstName //we can use p.firstname istead of (*p).firstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (s SLICE) updateMe(value string) {
	s[0] = value //ref type, hence the original slice will be updated

	// NOTE: type []string(slice) is of reference type

	/* BEHIND THE SCENE:
	Go have two datastructure, Array and slice
	-slice is what we use 99% of the time(unless a very good case scenario
	-slice datastructure will store, [lenght, capacity, ptr to head of array]
	-this "ptr to head" will have the ref to the underlying array
	-so, when we pass a slice into a function, it seems like we are doing a pass by ref(magically), but its not a magic, internally go compiler will just make a new copy of a slice(just like other datastrucre) but the value inside it the "ptr to head of the array" will not change, and hence we see a beautiful representation of pass by ref
	*/

}
