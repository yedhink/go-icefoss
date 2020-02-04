package main

import (
	"fmt"
	"os"
)

func main(){
	/*
	To loop over slices wecan make use of range function
	note that we can use _ if we have to avoid using
	a particular variable. literally temp values or discard values
	*/
	for _,arg := range os.Args {
		fmt.Print(arg,"  ")
	}
	/*
	Inorder to specify custom range, we can
	use python style slice index
	*/
	fmt.Println()
	for i,arg := range os.Args[1:] {
		fmt.Println(i," : ",arg)
	}
}
