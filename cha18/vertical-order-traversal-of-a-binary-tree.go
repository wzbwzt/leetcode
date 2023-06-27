package main

import "sort"

/*
给定二叉树，按垂序遍历返回其结点值。

对位于 (X, Y) 的每个结点而言，其左右子结点分别位于 (X-1, Y-1) 和 (X+1, Y-1)。

把一条垂线从 X = -infinity 移动到 X = +infinity ，每当该垂线与结点接触时，我们按从上到下的顺序报告结点的值（Y 坐标递减）。

如果两个结点位置相同，则首先报告的结点值较小。

按 X 坐标顺序返回非空报告的列表。每个报告都有一个结点值列表。



示例 1：



输入：[3,9,20,null,null,15,7]
输出：[[9],[3,15],[20],[7]]
解释：
在不丧失其普遍性的情况下，我们可以假设根结点位于 (0, 0)：
然后，值为 9 的结点出现在 (-1, -1)；
值为 3 和 15 的两个结点分别出现在 (0, 0) 和 (0, -2)；
值为 20 的结点出现在 (1, -1)；
值为 7 的结点出现在 (2, -2)。
示例 2：



输入：[1,2,3,4,5,6,7]
输出：[[4],[2],[1,5,6],[3],[7]]
解释：
根据给定的方案，值为 5 和 6 的两个结点出现在同一位置。
然而，在报告 "[1,5,6]" 中，结点值 5 排在前面，因为 5 小于 6。


提示：

树的结点数介于 1 和 1000 之间。
每个结点值介于 0 和 1000 之间。


*/

/*
思路：
`DFS`
1.对每个节点创建坐标位置

2.通过深度优先搜索, 定义两层哈希表:map[x]map[y][]int

3.对x，y,先后排序，获取对应[x][y] 下的节点值,追加到构建返回返回参数

复杂度:
时间复杂度： O(NlogN)，其中 N 为树的节点总数。
空间复杂度： O(N)，其中 N 为树的节点总数。
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) (out [][]int) {
	coords := make(map[int]map[int][]int, 0)

	dfs(root, 0, 0, coords)

	var keyX []int
	for x := range coords {
		keyX = append(keyX, x)
	}
	sort.Ints(keyX)

	for _, x := range keyX {
		var keyY []int
		for y := range coords[x] {
			keyY = append(keyY, y)
		}
		sort.Ints(keyY)

		nodev := []int{}
		for _, y := range keyY {
			sort.Ints(coords[x][y])
			nodev = append(nodev, coords[x][y]...)
		}

		out = append(out, nodev)
	}
	return out
}

func dfs(node *TreeNode, x int, y int, coords map[int]map[int][]int) {
	if node == nil {
		return
	}
	_, ok := coords[x]
	if !ok {
		coords[x] = map[int][]int{y: {}}
	}
	coords[x][y] = append(coords[x][y], node.Val)

	dfs(node.Left, x-1, y+1, coords)
	dfs(node.Right, x+1, y+1, coords)
}
