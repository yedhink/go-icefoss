package main

import "fmt"

func main() {
	/*
	copy, delete, insert
	*/
	var mainSlice []int = []int{1,2,3,4,5,6}
	// b := []uint{10: 1, 2} hows this even possible?
	// copy
	copySlice := make([]int,len(mainSlice))
	copy(copySlice,mainSlice)
	fmt.Println("copySlice \t: ",copySlice)
	fmt.Println("mainSlice \t: ",mainSlice)
	// delete - say n'th element
	n := 2
	deleteSlice := append(mainSlice[:n],mainSlice[(n+1):]...)
	// note that mainSlice is now changed since in append, dst is appended to
	fmt.Println("deleteSlice \t: ",deleteSlice)
	fmt.Println("mainSlice \t: ",mainSlice)
	fmt.Println("copySlice \t: ",copySlice)
	// insert - say n'th index
	n=1
	ourValue := []int{13}
	copySlice = append(copySlice[:n],append(ourValue,copySlice[n:]...)...)
	fmt.Println("copySlice \t: ",copySlice)
}
