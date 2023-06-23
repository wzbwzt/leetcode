package main

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

/*
思路：
`递归`
产品经理法



复杂度：
时间复杂度： O(N)
空间复杂度： O(N)

*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	num_left := maxDepth(root.Left)
	num_right := maxDepth(root.Right)
	if num_left < num_right {
		return num_right + 1
	} else {
		return num_left + 1
	}
}
