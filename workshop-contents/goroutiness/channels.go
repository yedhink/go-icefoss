package main

import "fmt"

func double(ip,op chan int){
	for {
		num := <-ip
		op <- num * 2
	}
}

func main() {
	to := make(chan int)
	from := make(chan int)
	go double(to,from)
	for i:=0;i<5;i++ {
		to <- i
		fmt.Printf("%d\t%d\n",i,<-from)
	}
}
