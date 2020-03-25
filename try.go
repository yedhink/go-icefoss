package main

import "fmt"

func main() {
	arr :=  [...]int{1,2,3,4,5}
	slice := arr[:3]
	s2 := []int{6,7}
	slic := append(slice,s2...)
	b := []uint{10:1, 2}
	fmt.Printf("Type: %T Value: %v\n", b,b)
	fmt.Printf("Type: %T Value: %v\n", slic,arr)
}
