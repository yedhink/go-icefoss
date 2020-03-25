package main

import "fmt"

func sqrt(number float64) float64 {
	var guess float64 = 1.0
	var count int = 1
	for guess*guess != number {
		guess -= ((guess*guess)-number)/(2*guess)
		if count++;count > 10 {
			break
		}
	}
	return guess
}

func main(){
	fmt.Printf("sqrt : %f\n",sqrt(2))
}
