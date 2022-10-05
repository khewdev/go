package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFlips(s string) int {
	n := len(s) // get length of original string
	s = s + s   // duplicates original string

	var target1, target2 string = "", ""

	for i := range s {
		if i%2 == 0 {
			target1 += "0"
		} else {
			target1 += "1"
		}

		if i%2 == 0 {
			target2 += "1"
		} else {
			target2 += "0"
		}
	}

	var res int

	var diff1, diff2, leftPtr int = 0, 0, 0

	for i := range s {
		if s[i] != target1[i] {
			diff1 += 1
		}

		if s[i] != target2[i] {
			diff2 += 1
		}

		if i-leftPtr+1 > n {
			if s[leftPtr] != target1[leftPtr] {
				diff1 -= 1
			}

			if s[leftPtr] != target2[leftPtr] {
				diff2 -= 1
			}

			leftPtr += 1
		}

		if i-leftPtr+1 == n {
			res = min(diff1, diff2)
		}
	}

	res = min(diff1, diff2)

	return res
}

func main() {
	fmt.Println("min flips", minFlips("111000"))
}
