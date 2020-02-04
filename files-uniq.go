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
		if dict[input.Text()] > 1{
			fmt.Printf("\t%s\t:%d\n",input.Text(),dict[input.Text()])
		}
	}
}

// bufio is buffered io
func main() {
	/*
	maps are the go variant of dictionaries
						key val
	*/
	counts := make(map[string]int)
	if len(os.Args[1:]) == 0{
		countLines(os.Stdin,counts)
	} else{
		file, err := os.Open("./files-uniq.test")
		if err != nil{
		fmt.Println("File doesnt exist")
		}
		fmt.Println("File was provided")
		countLines(file,counts)
	}
}

