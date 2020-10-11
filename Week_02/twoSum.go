func twoSum(nums []int, target int) []int {
	ret := []int{}
	if len(nums) < 2 {
		return ret
	}

	record := map[int]int{}
	for indx, val := range nums {
		min := target - val
		if suIndex, ok := record[min]; ok {
			ret = append(ret, indx)
			ret = append(ret, suIndex)
			return ret
		}
		record[val] = indx
	}

	return ret
}