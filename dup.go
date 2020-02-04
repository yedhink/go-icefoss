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
						key val
	*/
	dict := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		dict[input.Text()]++
		if dict[input.Text()] > 1{
			fmt.Print("\t",input.Text()," : ",dict[input.Text()],"\n")
		}
	}

	// LOOPING THROUGH DICTIONARY
	//for key,val := range dict {
	//	if val > 1{
	//	fmt.Println(key,"\t: ",val)
	//	}
	//}

}
