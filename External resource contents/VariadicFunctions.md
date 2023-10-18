### Why does Go not support overloading of methods and operators?
#### Official statement-
>Method dispatch is simplified if it doesn't need to do type matching as well. Experience with other languages told us that having a variety of methods with the same name but different signatures was occasionally useful but that it could also be confusing and fragile in practice. Matching only by name and requiring consistency in the types was a major simplifying decision in Go's type system.

## 2 tricks/ways to tackle the problem
by using either of the following-
- `Variadic`
- `...interface{}`

>we are already familiar with these functions: `fmt.Print()`, `fmt.Printf()` or `fmt.Println()`
> 
> They are nothing but Standard Golang *Variadic functions*
> 
> eg. where it is called with different parameter counts 
> ```go
> name := "Joe Bloe"
> age  := 23
> city := "Vancouver"
>
> // This call to fmt.Printf() has 4 parameters.
> fmt.Printf("%s is %d years old and lives in %s.\n", name, age, city)
>
> // This call to fmt.Printf() has 2 parameters.
> fmt.Printf("Hello %s!\n", name)
> ```

## Writing Our Own Variadic Functions and Methods
### 1. Variadic 

```go
package main

import "fmt"

// Sum returns the sum of adding all the numbers in
// the parameters together.
func Sum(numbers ...int) int {

    n := 0

    for _,number := range numbers {
        n += number
    }

    return n
}

func main() {
sm1 := Sum(1, 2, 3, 4) // = 1 + 2 + 3 + 4 = 10

    sm2 := Sum(1, 2) // = 1 + 2 = 3

    sm3 := Sum(7, 1, -2, 0, 18) // = 7 + 1 + -2 + 0 + 18 = 24

    sm4 := Sum() // = 0


    fmt.Printf("sm1 = %d\n", sm1)
    fmt.Printf("sm2 = %d\n", sm2)
    fmt.Printf("sm3 = %d\n", sm3)
    fmt.Printf("sm4 = %d\n", sm4)

    // Output:
    // sm1 = 10
    // sm2 = 3
    // sm3 = 24
    // sm4 = 0
}
```
the "real" type of numbers in our example ends up being `[]int`. (But of course, in our code we have to put `...int` and not `[]int` to tell Golang that this is a variadic function.)

#### Another example with string-
```go
package main

import "fmt"
import "github.com/reiver/go-stringcase"

// CamelCaseAll returns a []string that camelCases every string passed
// as a parameter.
func CamelCaseAll(ss ...string) []string {

    camelCasedStrings := make([]string, len(ss))

    for i,s := range ss {
        camelCasedStrings[i] = stringcase.ToCamelCase(s)
    }

    return camelCasedStrings
}

func main() {
ss1 := CamelCaseAll("hello world", "Apple Banana Cherry")
// = []string{"helloWorld", "appleBananaCherry"}

    ss2 := CamelCaseAll("More than meets the eye")
    // = []string{"moreThanMeetsTheEye"}

    ss3 := CamelCaseAll()
    // = []string{}


    fmt.Printf("ss1 = %v\n", ss1)
    fmt.Printf("ss2 = %v\n", ss2)
    fmt.Printf("ss3 = %v\n", ss3)

    // Output:
    // ss1 = [helloWorld appleBananaCherry]
    // ss2 = [moreThanMeetsTheEye]
    // ss3 = []
}
```

### 2. ...interface{}
>Limitation of variadic function with type ...int or ...string is that each argument passed to the variadic function must be of the same type.
>
> Solution is to use the variadic function with parameter type as `...interface{}`

```go
package main

import (
	"fmt"
	"reflect"
)

func Ul(things ...interface{}) {

	for _, it := range things {
		fmt.Println(it, "(", reflect.TypeOf(it), ")")
	}
}

func main() {
	// 2 parameter. 1st one a string. 2nd one an int.
	Ul("Hello world", 123)

	// 5 parameters. 1st, 3rd and 5th a string. 2nd a float64. 4th a int.
	Ul("apple", 7.2, "BANANA", 5, "cHeRy")

	// Output:
	//Hello world ( string )
	//123 ( int )
	//apple ( string )
	//7.2 ( float64 )
	//BANANA ( string )
	//5 ( int )
	//cHeRy ( string )
}
```

- But it is important to be aware of the cost to this.
We have lost compile time type checking of parameters.
In loosing that, and I think this is an important one, the defaulting and validation needs to be done at runtime.


- But regarless, if you need to do something like function overloading and method overloading in Golang, you can do it like this.

> **Resourse-** https://changelog.ca/log/2015/01/30/golang