package main

/*
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

*/

/*
思路：
`快慢指针`
前提：链表已经升序排序
根据高度平衡的二叉树原则，可以先选取中间节点作为根节点,左边的都为左节点，右边的都为右节点
对左右子链表再找中间节点,以此重复；
通过快慢指针找到中间节点



复杂度：
空间复杂度：空间复杂度为 O(logn)
时间复杂度：递归树的深度为 logn,每一层的基本操作数为 n,因此总的时间复杂度为O(nlogn)

*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	if head.Next.Next == nil {
		return &TreeNode{Val: head.Next.Val, Left: &TreeNode{Val: head.Val}}
	}

	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{Val: slow.Val}
	root.Right = sortedListToBST(slow.Next)

	pre.Next = nil
	root.Left = sortedListToBST(head)

	return root

}
