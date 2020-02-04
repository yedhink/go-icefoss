package main

// import can be multilined
// its traditional to write it in alphabetic order
import (
	"fmt"
	"os"
)

func main() {
	// is this implicit or explicit
	var args = os.Args
	i:=1
	for i < len(args) {
		fmt.Print(args[i],"  ")
		i++
	}

	// infinite loop without if condtion
	i--
	for {
		fmt.Print(args[i],"  ")
		i--
		if i<1{
			break
		}
	}
}

