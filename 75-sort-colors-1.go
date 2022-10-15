package main

func sortColors(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	start := 0
	end := len(nums) - 1
	index := 0

	for index <= end && start < end {
		if nums[index] == 0 {
			nums[index] = nums[start]
			nums[start] = 0
			start++
			index++
		} else if nums[index] == 2 {
			nums[index] = nums[end]
			nums[end] = 2
			end--
		} else {
			index++
		}
	}
}

func main() {
	input := []int{2, 0, 2, 1, 1, 0}
	sortColors(input)
}
