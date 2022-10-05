package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	first := strings.Split(s, "")
	second := strings.Split(t, "")

	sort.Strings(first)
	sort.Strings(second)

	return strings.Join(first, "") == strings.Join(second, "")
}

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Print(isAnagram(s, t))
}
