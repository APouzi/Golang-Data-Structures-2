package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	x := primes[3:]
	s = append(s, 1, 2)
	fmt.Println(s, x)

	// Note: You can't actually delete from the slices, you would have to do some "hacky" things. An example: https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
}
