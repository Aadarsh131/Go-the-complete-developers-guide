package main

import (
	"fmt"
	"reflect"
)

func variadicFunc(param ...interface{}) {
	fmt.Println(reflect.TypeOf(param[0]))
}

func main() {
	variadicFunc(45)
}
