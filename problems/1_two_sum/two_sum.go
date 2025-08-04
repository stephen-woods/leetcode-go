// Package problem1
//
// Given an array of integers nums and an integer target, return indices of the
// two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may
// not use the same element twice.
//
// You can return the answer in any order.
package problem1

func twoSum(nums []int, target int) []int {

	indices := map[int]int{}

	for i, num := range nums {
	  diff := target - num	

		j, exists := indices[diff]
		if exists {
			return []int {i, j}
		} else {
			indices[num] = i
		}
	}
	panic("No two sum solution found") // This should never happen given the problem constraints
}
