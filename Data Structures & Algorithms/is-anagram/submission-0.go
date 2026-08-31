func isAnagram(s string, t string) bool {
	count := make(map[rune]int)
	for _, char := range s{
		count[char]++
	}

	for _, char := range t{
		count[char]--
	}

	for _,val:=range count{
		if val!=0{
			return false
		}
	}
	return true
}
