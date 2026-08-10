package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		println(i, num)
	}
	fmt.Println("   ")
	for i := 0; i < len(nums); i++ {
		println(i, nums[i])
	}
}
