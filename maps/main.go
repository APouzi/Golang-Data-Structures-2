package main

import "fmt"

func main() {

	hello := map[string]string{}
	hello["test"] = "yes"
	// IF statements in Go can include both a condition and intialization statment
	if val, ok := hello[""]; ok {
		fmt.Println("membership in", val)
	} else {
		fmt.Println("membership is", ok)
	}

}
