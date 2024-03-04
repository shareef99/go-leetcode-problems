package main

import (
	"fmt"
)

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

// Example 1:
// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].

// Example 2:
// Input: nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]

// Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?

// Link: https://leetcode.com/problems/squares-of-a-sorted-array/

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(optimizeSortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(nums []int) []int {
	squareNumber := []int{}

	for i := 0; i < len(nums); i++ {
		squareNumber = append(squareNumber, nums[i]*nums[i])
	}

	for i := 0; i < len(squareNumber); i++ {
		for j := i + 1; j < len(squareNumber); j++ {
			if squareNumber[i] > squareNumber[j] {
				squareNumber[i], squareNumber[j] = squareNumber[j], squareNumber[i]
			}
		}
	}

	return squareNumber
}

func optimizeSortedSquares(nums []int) []int {
	squareNumber := make([]int, len(nums))

	left := 0
	right := len(nums) - 1
	index := len(nums) - 1

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			squareNumber[index] = leftSquare
			left++
		} else {
			squareNumber[index] = rightSquare
			right--
		}
		index--
	}

	return squareNumber
}
