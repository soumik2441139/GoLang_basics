package main

import "fmt"

func main() {
	//unintialized slice is nil
	//var nums []int

	//fmt.Println(nums == nil)
	//fmt.Println(len(nums))

	/*var nums = make([]int, 2, 5)
	fmt.Println(nums == nil)
	fmt.Println(cap(nums)) //capacity :max number of elements can fit
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums[0] = 10
	nums[1] = 20
	fmt.Println(nums)
	fmt.Println(cap(nums)) //capacity :max number of elements can fit*/

	/*var nums = make([]int, 0, 5)
	nums = append(nums, 1)

	var nums2 = make([]int, len(nums))
	copy(nums2, nums)

	fmt.Println(nums, nums2)*/

	/*var nums = []int{1, 2, 3, 4, 5}
	fmt.Println(nums[1:3])
	fmt.Println(nums[:2])
	fmt.Println(nums[2:])*/

	/*var nums = []int{1, 2}
	var nums2 = []int{1, 2}

	fmt.Println(slices.Equal(nums, nums2))*/

	var nums = [][]int{{1, 2}, {3, 4}}
	fmt.Println(nums)

}
