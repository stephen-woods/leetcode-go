package problem2

import (
	"testing"
)

// Example 1:
//
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
func TestAddTwoNumbers1(t *testing.T) {
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})

	result := addTwoNumbers(l1, l2)

	expected := createList([]int{7, 0, 8})

	if !ListNodeEquals(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Example 2:
//
// Input: l1 = [0], l2 = [0]
// Output: [0]
func TestAddTwoNumbers2(t *testing.T) {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}

	result := addTwoNumbers(l1, l2)

	expected := createList([]int{0})

	if !ListNodeEquals(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Example 3:
//
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]
func TestAddTwoNumbers3(t *testing.T) {
	l1 := createList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := createList([]int{9, 9, 9, 9})

	result := addTwoNumbers(l1, l2)

	expected := createList([]int{8, 9, 9, 9, 0, 0, 0, 1})

	if !ListNodeEquals(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Creates a linked list from a slice of integers.
func createList(nums []int) *ListNode {
	var ret *ListNode = nil

	for i := len(nums) - 1; i >= 0; i-- {
		n := &ListNode{
			Val:  nums[i],
			Next: ret,
		}
		ret = n
	}
	return ret
}

