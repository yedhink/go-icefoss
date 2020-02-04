package main

// import can be multilined
// its traditional to write it in alphabetic order
import (
	"bufio"
	"fmt"
	"os"
)

// bufio is buffered io
func main() {
	/*
	maps are the go variant of dictionaries
	*/
	dict := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		dict[input.Text()]++
	}

	for key,val := range dict {
		fmt.Println(key,"\t: ",val)
	}

}
