package main

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例 1：
```
![image](https://p.ipic.vip/r1afvu.jpg)

```
输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]


提示：

链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100

*/

/*
思路：
A->B->C->D
技巧：对于一个head node可能改变的链表，可以通过虚拟一个空节点来指代；
两两交换时涉及到四个节点，prehead->A->B->C

交换步骤：
1. prehead->A->C
2. B->A
3. prehaed->B

迭代交换


复杂度：
时间复杂度: O(n)
空间复杂度：O(1)

*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prehead := &ListNode{}
	prehead.Next = head

	pre := prehead
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next

		pre = cur
		cur = cur.Next
	}

	return prehead.Next
}

/*
思路：
递归

复杂度：
时间复杂度: O(n)
空间复杂度：O(1)


*/

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs2(next.Next)
	next.Next = head
	return next
}
