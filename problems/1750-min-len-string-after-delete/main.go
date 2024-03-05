package main

import "fmt"

func main() {
	fmt.Println("cs", minimumLength("ca"))               // Output: 2
	fmt.Println("cabaabac", minimumLength("cabaabac"))   // Output: 0
	fmt.Println("aabccabba", minimumLength("aabccabba")) // Output: 3
}

func minimumLength(s string) int {
	left, right := 0, len(s)-1

	for left < right {
		lc := s[left]
		rc := s[right]

		if lc != rc {
			return right - left + 1
		}

		for left+1 < right && lc == s[left+1] {
			left++
		}

		for left < right-1 && rc == s[right-1] {
			right--
		}

		right--
		left++
	}

	return right - left + 1
}
