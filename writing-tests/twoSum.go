package twoSum


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)


	for i, n := range nums {
		term := target - n
		index, check := m[term]
		if (check) {
			return []int{index, i}
		}
		m[n] = i
	}
	return nil
}