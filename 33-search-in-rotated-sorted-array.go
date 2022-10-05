package main

import "fmt"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		fmt.Println("l", l)
		fmt.Println("r", r)
		fmt.Println("mid", mid)
		fmt.Println("nums[mid]", nums[mid])
		fmt.Println("nums[l]", nums[l])

		if target == nums[mid] {
			fmt.Println("target == nums[mid]")
			return mid
		}

		if nums[l] <= nums[mid] {
			fmt.Println("nums[l] <= nums[mid]")
			if target > nums[mid] || target < nums[l] {
				fmt.Println("target > nums[mid] || target < nums[l]")
				l = mid + 1
			} else {
				fmt.Println("target > nums[mid] || target < nums[l] else")
				r = mid - 1
			}
		} else {
			fmt.Println("nums[l] <= nums[mid] else")
			if target < nums[mid] || target > nums[r] {
				fmt.Println("target < nums[mid] || target > nums[r]")
				r = mid - 1
			} else {
				fmt.Println("target < nums[mid] || target > nums[r] else")
				l = mid + 1
			}
		}
		fmt.Println()
	}

	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}
