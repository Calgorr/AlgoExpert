package main

import "fmt"

func main() {
	nums := []int{2, 1, 2, 2, 2, 3, 4, 2}
	fmt.Println(Move(nums, 2))
}

func Move(nums []int, toMove int) []int {
	first, last := 0, len(nums)-1
	for {
		if first >= last {
			break
		}
		if nums[first] != 2 {
			first++
		}
		if nums[last] == 2 {
			last--
		}

		if nums[last] != 2 && nums[first] == 2 {
			nums[last], nums[first] = nums[first], nums[last]
		}

	}
	return nums
}
