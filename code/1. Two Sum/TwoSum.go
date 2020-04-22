package src

func TwoSum(nums []int, target int) []int {
	s := map[int]int{}
	for i, v := range nums {
		if v, ok := s[target-v]; ok {
			return []int{i, v}
		}
		s[v] = i
	}
	return []int{}
}