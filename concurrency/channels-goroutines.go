package main
import "fmt"

func routine(s string,c chan string){
	fmt.Println("Process done inside routine : ",s)
	c <- s + "done"
}

func square(ip chan int,op chan int) {
	for {
		num := <-ip
		op <- num*num
	}
}

func fibo(n int,c chan int) {
	prev,next := 0, 1
	for i:=0;i<n;i++ {
		c<-prev
		prev, next = next, prev+next
	}
	close(c)
}

func main(){
	c1 := make(chan int)
	c2 := make(chan int)
	go square(c1,c2)
	for i:=0;i<10;i++ {
		c1<-i
		fmt.Println(i," : ",<-c2)
	}
	go fibo(10,c2)
	for i := range c2 {
		fmt.Println("fibonacci : ",i)
	}
}
