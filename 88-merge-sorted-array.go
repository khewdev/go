package main

import "fmt"

func main() {
	m := 3
	n := 3
	nums1 := [6]int{3, 3, 3}
	nums2 := [3]int{1, 2, 6}

	fmt.Println(nums1)
	fmt.Println(nums2)

	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
		fmt.Println(nums1)
	}

	fmt.Println(nums1)

	fmt.Println(n)

	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}

	fmt.Println(nums1)
}
