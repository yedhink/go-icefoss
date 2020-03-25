package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func main() {
	url := "http://nibrahim.net.in/"
	resp,err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr,"ERORR",err)
		os.Exit(1)
	}
	b,err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s",b)
}
