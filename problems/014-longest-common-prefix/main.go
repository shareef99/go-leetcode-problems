package main

import "fmt"

// 14. Longest Common Prefix
// Easy
// Topics
// Companies
// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

func main() {
	fmt.Println(longestCommonPrefix([]string{"flow", "flower", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	var prefix string
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return prefix
			}
		}
		prefix = prefix + string(strs[0][i])
	}
	return prefix
}
