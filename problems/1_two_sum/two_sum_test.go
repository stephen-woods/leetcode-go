package problem1

import (
	"leetcode-go/utils"
	"testing"
)

// Example 1:
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	expected := []int{0, 1}

	if !utils.EqualUnordered(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}


// Example 2:
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6 

	result := twoSum(nums, target)

	expected := []int{1, 2}

	if !utils.EqualUnordered(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}


// Example 3:
//
// Input: nums = [3,3], target = 6
// Output: [0,1]
func TestTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6 

	result := twoSum(nums, target)

	expected := []int{0, 1}

	if !utils.EqualUnordered(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
