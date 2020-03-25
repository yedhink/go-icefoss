# Basics
* all programs are part of package
* we can import code in a factored manner
* variables declared locally within a function should be used somewhere, else would results in error
### exported names
* in go all exported names start with capital letter
* from package we can access only the exported names
### go reverse syntax for declaration
* the idea is that, its much more easier to read from left to right rather than the other way around like in C
* its' clear when we write a function which itself is a ptr and includes function ptrs as its args and returns a function itself
	```go
	outer func(x,y int) (inner func(z int) *int, name string) *int
	```
* functions can return N number of items
#### Naked returns
```
package main

import "fmt"

func naked(arg int) (x,y int) {
	x = arg/2 
	y = x-arg
	return
}

func main() {
	var1,var2 := naked(10)
}
```
### Factoring
including set of statements within in parens is called factoring. eg we do multiple imports using factors. Also multiple variable declaration, like so:-
	```go
var (
x int = 10
y bool = true
z string = "asdas"
	)
	```

# control-flow
## for
	```go
	for i:=0;i<10;i++ {
		fmt.Printf(i)
	}

	i:=1
	for  ;i < 10; {
		i+=1
	}

	for  i < 10 {
		i+=1
	}

	for {
		i*=2
	}
	```
## if-else statement
	```go
	if i < 10 {

	}

	if v:=10;  v%2==0 {
		// v only lives within this scope
	}

	if v:=10;  v%2==0 {
		// v only lives within this scope
	} else {
		// v is  also accessible here
	}
	```

# Arrays
* creating zero value array of size n
```go
var arr [5]int;
```
* if element types of two arrays are same, then we can compare it
* dynamic length can be attained using ellipsis

# Slices
* slices are like references to an array
* slice is a variadic function
* it's a struct of sorts :- 
```go
	type slice struct {
		zerothElement *type
		len int
		cap int
	}
```
* they don't have a fixed length, rather they can grow in size
* using make we can create dynamically sized arrays :- `make([]int,len,cap)`
