func twoSum(nums []int, target int) []int {
    comp := make(map[int]int)

	for i, num := range nums{
		diff := target-num
		idx, exists := comp[diff]
		if exists{
			return []int{idx,i}
		}
		comp[num]=i
	}
	return []int{}
}
