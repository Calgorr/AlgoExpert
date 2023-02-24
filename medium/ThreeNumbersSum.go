package main

import "fmt"

func main() {
	array := []int{5, 3, -8, 2, -6, 12, 1, 6}
	fmt.Println(ThreeSum(array, 0))
}

func ThreeSum(nums []int, target int) [][3]int {
	var res [][3]int
	for i := 0; i < len(nums); i++ {
		memory := make(map[int]int)
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if _, ok := memory[target-nums[i]-nums[j]]; ok {
				res = append(res, [...]int{nums[i], nums[j], target - nums[i] - nums[j]})
			} else {
				memory[nums[j]] = j
			}
		}
	}
	return res
}
