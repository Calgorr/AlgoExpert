package TwoNumbersSum

func PairSum(nums []int, sum int) [2]int {
	Memory := make(map[int]struct{})
	for _, v := range nums {
		num := sum - v
		if _, ok := Memory[v]; ok {
			return [2]int{num, v}
		} else {
			Memory[num] = struct{}{}
		}
	}
	return [...]int{0, 0}
}
