func groupAnagrams(strs []string) [][]string {
	count := make(map[[26]int][]string)

	for _, str := range strs {
		arr := [26]int{}
		for _, char := range str {
			idx := char - 'a'
			arr[idx] += 1
		}

		count[arr] = append(count[arr], str)
	}

	res := [][]string{}
	for _, val := range count {
		if len(val) > 0 {
			res = append(res, val)
		}
	}
	return res
}