package main

import "fmt"

func slidingWindow(arr []int, k int) {
	bank := []int{}
	limit := len(arr) - k
	for i := 0; i < limit+1; i++ {
		for j := i; j < k+i; j++ {
			bank = append(bank, arr[j])
		}
	}
	fmt.Println(bank)
}

//Largest sum in the K window     1,2,4,1,2,4,5,1   k = 3
func pointerWindowTrailing(arr []int, k int) {
	var ans = []int{}
	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		fmt.Print(sum)
		if i >= (k - 1) {
			ans = append(ans, sum)
			sum -= arr[i-(k-1)]
		}
	}
	fmt.Println(ans)

}

func PointerWindowTrailingPrac(nums []int, k int) {
	var sum int = 0
	var numsAns []int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i >= (k - 1) {
			numsAns = append(numsAns, sum)
			sum -= nums[i-(k-1)]
		}
	}
	fmt.Println(numsAns)
}

// func pointerWindow(arr []int, k int) {
// 	var a int = 0
// 	var b int = 0
// 	for i := 0; i < len(arr)-1; i++ {
// 		if arr[i] < arr[i+1] {
// 			b = i + 1

// 		} else {
// 			a = i + 1
// 		}

// 	}
// }

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	arr2 := []int{1, 2, 4, 1, 2, 4, 5, 1}
	slidingWindow(arr, 3)
	pointerWindowTrailing(arr2, 3)
	PointerWindowTrailingPrac(arr, 3)

	// fmt.Print("test")

}
