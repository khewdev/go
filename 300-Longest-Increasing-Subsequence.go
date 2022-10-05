package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	lis := []int{}

	for _, num := range nums {
		index := sort.SearchInts(lis, num)

		if index == len(lis) {
			lis = append(lis, num)
		} else if num < lis[index] {
			lis[index] = num
		}
	}

	return len(lis)
}

func main() {
	nums := []int{0, 1, 0, 3, 2, 3}

	fmt.Println(lengthOfLIS(nums))
}
