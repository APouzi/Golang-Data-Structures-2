package main

import "fmt"

func subArrayAdding(k int, arr []int) {
	length := len(arr) - k
	ans := []int{}
	for i := 0; i < length+1; i++ {
		sum := 0
		for j := i; j < (length + i); j++ {
			sum += arr[j]
		}
		ans = append(ans, sum)
	}
	fmt.Println(ans)
}

func main() {
	var repeat = [...]int{1, 2, 3, 4, 4, 4, 5, 5, 5, 5, 6}
	fmt.Println(pointerRepeat(repeat[:]))
	// var array1 = [6]int{1, 2, 3, 4, 5, 6}
	// subArrayAdding(3, array1[:])
	// var array2 = [...]int{1, 2, 3, 1, 2, 3, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3}
	// subPointer(array2[:])

}

func subPointer(arr []int) {
	var a int = 0
	var b int = 0
	var count int = 1
	var bank int = 0

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			count++
			b = i + 1
		} else {
			b = 0
			a = i + 1
			count = 1
		}
		if count >= bank {
			bank = count
		}
	}
	slice := arr[a : b+1]
	fmt.Println(slice, bank, count)
}

func pointerRepeat(arr []int) int {
	a := 0
	b := 0
	max := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			b = i
			if (b-a)+1 > max {
				max = (b - a) + 1
			}

		} else {
			a = i
		}

	}
	return max
}
