package leetcode

// You are given two non-empty linked lists representing two non-negative
// integers. The digits are stored in reverse order, and each of their nodes
// contains a single digit. Add the two numbers and return the sum as a linked
// list.
//
// You may assume the two numbers do not contain any leading zero, except the
// number 0 itself.

type ListNode struct {
	Val  int
	Next *ListNode
}

// ListNodeEquals checks if two linked lists are equal.
func ListNodeEquals(ln1, ln2 *ListNode) bool {
	for ln1 != nil && ln2 != nil {
		if ln1.Val != ln2.Val {
			return false
		}
		ln1 = ln1.Next
		ln2 = ln2.Next
	}
	return ln1 == nil && ln2 == nil
}

// addTwoNumbers adds two numbers represented by linked lists and returns the
// result as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var sum int
	var carry int

	// Initial node
	sum = l1.Val + l2.Val
	carry = sum / 10
	sum = sum % 10
	
	head := &ListNode{
		Val:  sum,
		Next: nil,
	}

	// Remaining nodes
	p := head
	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil || l2 != nil {
		sum = 0
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		sum += carry
		carry = sum / 10
		sum = sum % 10

		p.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}

		// Move pointers
		p = p.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	// Check to see if there is a carry left
	if carry > 0 {
		p.Next = &ListNode{
			Val: carry,
		}
	}

	return head
}
