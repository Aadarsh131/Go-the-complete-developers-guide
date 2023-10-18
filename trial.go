package main

import (
	"fmt"
)

type stringSliceType []string

func main() {
	a := stringSliceType{}
	a.someFunc(stringSliceType{"a", "b"})
}

func (stringSliceType) someFunc(strings []string) {
	fmt.Println(len(strings))
}
