package main

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

*/

/*
思路：
`快慢指针`
eg. A -> B -> C -> D -> E 右移 2 位 D-E-A-B-C
原则：不管移动多少，节点之间的相对位置是不变的
规则：移动k位，倒数第k位会移动到首位,即D会移动到head，倒数第K+1位会移动到last
注意：移动k位和移动k%len效果等同
过程：通过快慢指针回去倒数第K+1个节点


复杂度：
时间复杂度：节点最多只遍历两遍，时间复杂度为 O(n)
空间复杂度：未使用额外的空间，空间复杂度 O(1)
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := 0
	c := head
	for c != nil {
		length++
		c = c.Next
	}
	k = k % length
	if k == 0 {
		return head
	}
	var slow, fast = head, head
	for i := 0; i < k+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	next := slow.Next
	slow.Next = nil
	out := next
	for next.Next != nil {
		next = next.Next
	}
	next.Next = head

	return out
}
