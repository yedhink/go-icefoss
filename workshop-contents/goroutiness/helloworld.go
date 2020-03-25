package main
import (
"fmt"
"time"
)

func repeat(s string){
	for i:=0;i<5;i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	// literally runs the repeat as a seperate thread
	go repeat("world")
	repeat("hello")
}
