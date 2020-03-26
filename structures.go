package main
import "fmt"

type Student struct {
	name string
	age int
}

type Class struct {
	students map[string]Student
}

func main(){
	s1 := Student {
		age : 21,
		name : "Yedhin",
	}
	s2 := Student{"Zach",21}
	s2.age = 22
	fmt.Println(s1,s2)

	c1 := Class{
		students : map[string]Student {
			"FIT16CS125" : Student{
				age : 21,
				name : "Yedhin",
			},
		},
	}
	fmt.Println(c1)
}
