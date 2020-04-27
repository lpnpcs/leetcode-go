package src



func lengthOfLongestSubstring(s string) int {
	index := [256] int{}
	ans  := 0
	i:=0
	for j:=0 ;j < len(s); j++{
		if index[s[j]] > 0 && index[s[j]] > i {
			i = index[s[j]]
		}
		index[s[j]] = j+1
		ans = maxNums(ans, j-i + 1)

	}
	return ans
}

func maxNums(a, b int) int {
	if a > b {
		return a
	}
	return b
}

