package main
import "fmt"

type Student struct {
	name string
	age int
}

func main(){
	// literal map
	var literalMap = map[string]Student {
		"FIT16CS125":Student {
			"Yedhin",21,
		},
		"FIT16CS126":Student {
			"Zach",21,
		},
	}
	for key,value := range literalMap {
		fmt.Printf("%v : %v\n",key,value)
	}

	// non-literal
	nonLiteralMap := make(map[string]Student,2)
	nonLiteralMap["FIT16CS126"] = Student{
		"Yedhin",21,
	}
	fmt.Println(nonLiteralMap)

	// insert, delete, search
	// insert
	literalMap["newItem"] = Student{"NewStud", 18}
	fmt.Println(literalMap)
	// delete
	delete(literalMap,"newItem")
	fmt.Println(literalMap)
	// search
	value,exists := literalMap["newItem"]
	fmt.Printf("value for non-exist -> value : %v and exist : %v",value,exists)
	value,exists = literalMap["FIT16CS125"]
	fmt.Printf("value for existing key -> value : %v and exist : %v",value,exists)
}
