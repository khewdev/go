package main

import "fmt"

type CharCount [26]int

func groupAnagrams(strs []string) [][]string {
	hash := make(map[CharCount][]string)
	for _, word := range strs {
		charCount := CharCount{}
		for i := 0; i < len(word); i++ {
			charCount[word[i]-'a']++
		}
		// fmt.Println(charCount)

		hash[charCount] = append(hash[charCount], word)
	}

	fmt.Println(hash)

	var result [][]string
	for _, value := range hash {
		result = append(result, value)
	}

	return result
}

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(str))
}
