package main
import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	go routine("goroutine 1")
	routine("mainroutine1")
	go routine("goroutine 2")
	fmt.Scanln()
}

func routine(s string){
	for i:=0;i<3;i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
