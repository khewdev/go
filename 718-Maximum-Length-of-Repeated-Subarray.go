package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLength(nums1 []int, nums2 []int) int {
	// if len(nums1) > len(nums2) {
	//     nums1, nums2 = nums2, nums1
	// }

	dp := make([][]int, len(nums1)+1)

	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	result := int(0)

	for i := 1; i < len(nums1)+1; i++ {
		for j := 1; j < len(nums2)+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			result = max(result, dp[i][j])
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4}

	fmt.Println(findLength(nums1, nums2))
}
