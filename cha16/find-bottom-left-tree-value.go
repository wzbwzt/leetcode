package main

/*
给定一个二叉树，在树的最后一行找到最左边的值。

示例 1:

输入:

    2
   / \
  1   3

输出:
1


示例 2:

输入:

        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7

输出:
7



*/

/*
 思路：
 `BFS`
广度优先搜索:创建队列，放入待处理节点，创建数组记录已处理节点


复杂度：
时间复杂度： O(N), N 为树的节点数
空间复杂度： O(Q)，其中 Q 为队列长度，最坏的情况是满二叉树，此时和 N 同阶，其中 N 为树的节点总数

*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var out int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		out = queue[0].Val

		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}

	return out

}

/*
 思路：
 `DFS`
深度优先搜索: 它从起始节点开始，沿着一条路径尽可能深地探索，直到到达最深处或无法继续前进时回溯，然后继续探索其他路径。DFS 可以用递归或栈来实现。


复杂度：
时间复杂度： O(N), N 为树的节点数
空间复杂度： O(h), h 为树的高度

*/

// Definition for a binary tree node.

func findBottomLeftValue2(root *TreeNode) int {
	var maxdepth int
	var out int = root.Val

	dfs(root, 0, &out, &maxdepth)

	return out

}

func dfs(node *TreeNode, depth int, out *int, maxdepth *int) {
	if node == nil {
		return
	}

	if depth > *maxdepth {
		*out = node.Val
		*maxdepth = depth
	}

	dfs(node.Left, depth+1, out, maxdepth)
	dfs(node.Right, depth+1, out, maxdepth)

}
