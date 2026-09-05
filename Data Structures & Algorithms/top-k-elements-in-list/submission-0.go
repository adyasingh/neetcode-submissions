func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums{
		count[num]++
	}

	res := []int{}
	for i:=0;i<k;i++{
		maxVal:=0
		maxKey:=0
		for key, val:= range count{
			if val>maxVal{
				maxVal=val
				maxKey=key
			}
		}

		res = append(res, maxKey)
		delete(count, maxKey)
	}
	return res

}
