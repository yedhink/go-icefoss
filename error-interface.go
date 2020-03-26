package main
import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

func (err MyError) Error() string {
	return fmt.Sprintf("%v\t: %v",err.when,err.what)
}

func Sqrt(v float64) (float64,error) {
	if v < 0 {
		return v, MyError{time.Now(),"You've Entered a negative number"}
	}
	x0 := 1.0
	count := 0
	for x0 != v {
		if count++;count > 10 {
			break
		}
		x0 -= ((x0*x0)-v)/(2*x0)
	}
	return x0,nil
}

func main(){
	fmt.Println(Sqrt(10.0))
	fmt.Println(Sqrt(-10.0))
}
