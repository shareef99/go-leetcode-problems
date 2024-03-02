package main

import (
	"fmt"
)

// Description:
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

// Link: https://leetcode.com/problems/two-sum/description/

func main() {
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(optimizeTwoSum([]int{3,2,4}, 6))
}

func twoSum(nums []int, target int) [2]int {
	result := [2]int{}

	for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums); j++ {
					if i == j {
							continue;
					}

					if nums[i] + nums[j] == target {
							result[1] = i;
							result[0] = j;
					}
			}
	}

	return result;
}

func optimizeTwoSum(nums []int, target int) [2]int {
	result := [2]int{}

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
			m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
			if _, ok := m[target - nums[i]]; ok && m[target - nums[i]] != i {
					result[0] = i
					result[1] = m[target - nums[i]]
					break
			}
	}

	return result
}