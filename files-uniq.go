package main

// import can be multilined
// its traditional to write it in alphabetic order
import (
	"bufio"
	"fmt"
	"os"
)

func countLines(file *os.File,dict map[string]int){
	input := bufio.NewScanner(file)
	for input.Scan() {
		dict[input.Text()]++
	}
}

// bufio is buffered io
func main() {
	counts := make(map[string]int)
	// inorder to check type of a variable use this
	fmt.Printf("Type of os.Stdin = %T\n",os.Stdin)
	if len(os.Args[1:]) == 0{
		countLines(os.Stdin,counts)
	} else{
		for  _,item := range os.Args[1:]{
			file, err := os.Open(item)
			if err != nil{
				fmt.Println(err)
			}
			countLines(file,counts)
		}
	}
	for key,val := range counts {
		fmt.Printf("%s\t: %d\n",key,val)
	}
}
