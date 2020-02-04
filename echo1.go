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
	//args := os.Args
	// the open brace should be on same line
	// :=  is implicit declaration
	// os.Args is a slice of strigs or a list of strings
	for i:=1; i < len(args);i++ {
		fmt.Print(args[i]," ")
	}
}
