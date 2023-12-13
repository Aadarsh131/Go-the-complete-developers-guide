package main

import "fmt"

type sameNamedStudents func(int, string)
type FUNCTION func(int, int) int

func (s sameNamedStudents) countNames() {
	s(4, "Atul")
	fmt.Println("counting total duplicate names")
}

func main() {
	a := sameNamedStudents(func(count int, name string) {
		fmt.Println("There are", count, name, "named students")
	})
	// a(3, "Aadarsh")
	a.countNames()
	fmt.Println("-------------")

	var b sameNamedStudents = func(i int, s string) {
		fmt.Println(i, s)
	}
	b.countNames()

	var add FUNCTION = func(x int, y int) int {
		return x+y
	}
	fmt.Println(add(4,5))
}
