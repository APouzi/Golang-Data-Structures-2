package main

import "fmt"

func main() {
	var words [7]string
	words[0] = "hello"
	words[1] = "hello1"
	words[2] = "hello2"
	words[3] = "hello3"
	words[4] = "hello4"
	words[5] = "hello5"
	words[6] = "hello6"
	wordsSlice := words[0:4]
	fmt.Println(wordsSlice)

	array := [3]string{}
	array[0] = "hello"
	array[1] = "hello"
	array[2] = "hello"
	fmt.Println(array)
}
