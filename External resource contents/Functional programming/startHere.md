# Function as type
```golang
type FUNCTION func(int,int) int
```

```golang
func main(){
    var add FUNCTION = func(x int, y int) int{
        return x+y
    }
    result := add(3,4) //7
}
```
OR,

```golang
func main(){
    add := FUNCTION(func(x int, y int) int{
        return x+y
    })
    result := add(3,4) //7
}
```
[Another Example](./codes/funcAsTypes.go)

# Function as parameter

```golang
package main

import "fmt"

func wrapper(add func(int, int) int) {
	fmt.Println(add(4, 5))
}

func doer(a int, b int) int {
	return a + b
}

func main() {
	wrapper(doer)
}
```
