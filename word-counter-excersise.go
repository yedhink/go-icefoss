/*
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/
package main
import(
	"fmt"
	"strings"
)

func getWordCount(s string) map[string]int {
	wordCountMap := make(map[string]int)
	word := ""
	for _,v := range s{
		if v == 32 {
			wordCountMap[word] += 1
			word = ""
			continue
		}
		word = word + fmt.Sprintf("%s",string(v))
	}
	wordCountMap[word] += 1
	return wordCountMap
}

func WordCount(s string) map[string]int {
	var stringWordCount = make(map[string]int)
	for _,word := range strings.Fields(s){
		fmt.Println(word)
		stringWordCount[word] += 1
	}
	return stringWordCount
}

func main(){
	var s string = "hi hello how are hi"
	wordCountMap := WordCount(s)
	fmt.Println(wordCountMap)
}
