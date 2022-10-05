package main

import (
	"fmt"
	"math"
	"sort"
)

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)

	fmt.Println(nums)

	res := 0
	mod := math.Pow(10, 9) + 7

	fmt.Println(mod)

	r := len(nums) - 1

	fmt.Println(r)

	for i, left := range nums {
		for left+nums[r] > target && i <= r {
			r = r - 1
		}

		if i <= r {
			res = int(float64(res) + math.Pow(2, float64(r)-1))
			res = int(math.Mod(float64(res), mod))
		}
	}

	return res
}

func main() {
	nums := []int{3, 5, 6, 7}
	fmt.Println(nums)
	target := 9

	fmt.Println(numSubseq(nums, target))

}
