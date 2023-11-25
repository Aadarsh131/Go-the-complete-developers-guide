package main

import (
	"fmt"
)

func RemoveIndex(s []int, index int) {
	append(s[:index], s[index+1:])
	fmt.Println(s)
}

func main() {
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("all: ", all) //[0 1 2 3 4 5 6 7 8 9]
	RemoveIndex(all, 5)
	fmt.Println("all: ", all) //[0 1 2 3 4 6 7 8 9 9]

	// fmt.Println("removeIndex: ", removeIndex) //[0 1 2 3 4 6 7 8 9]

	// removeIndex[0] = 999
	// fmt.Println("all: ", all) //[999 1 2 3 4 6 7 9 9]
	// fmt.Println("removeIndex: ", removeIndex) //[999 1 2 3 4 6 7 8 9]
}
