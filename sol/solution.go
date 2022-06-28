package sol

func rob(nums []int) int {
	numsLen := len(nums)
	prevTwo := 0
	prevOne := nums[numsLen-1]
	currentMax := 0
	for start := numsLen - 2; start >= 0; start-- {
		currentMax = max(nums[start]+prevTwo, prevOne)
		prevTwo = prevOne
		prevOne = currentMax
	}
	return max(prevOne, prevTwo)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
