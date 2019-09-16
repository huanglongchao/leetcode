package main

/**
	Example:

	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 0 -> 8
	Explanation: 342 + 465 = 807.
 */

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0,nil}
	var p,q,cur = l1,l2,result

	carry := 0 //进位
	for p!=nil || q != nil{
		x,y := 0,0
		if p!=nil{x = p.Val}
		if q!=nil{y = q.Val}
		sum := x+y+carry
		carry = sum / 10
		cur.Next = &ListNode{sum%10,nil}
		cur = cur.Next
		if p!=nil{p = p.Next}
		if q!=nil{q= q.Next}
	}
	if carry > 0 {
		cur.Next = &ListNode{carry,nil}
	}
	return result.Next
}
