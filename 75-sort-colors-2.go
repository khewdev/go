package main

import "fmt"

func sortColors(nums []int) {
	swapCount := 1

	for swapCount > 0 {
		swapCount = 0
		for itemIndex := 0; itemIndex < len(nums)-1; itemIndex++ {
			if nums[itemIndex] > nums[itemIndex+1] {
				nums[itemIndex], nums[itemIndex+1] = nums[itemIndex+1], nums[itemIndex]
				swapCount += 1
				fmt.Printf("%v\n", nums)
			}
		}
		fmt.Println("swapCount", swapCount)
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
}
