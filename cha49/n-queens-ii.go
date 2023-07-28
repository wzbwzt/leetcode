package main

/*
https://leetcode.cn/problems/n-queens-ii

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。


示例 1：
```

![](https://p.ipic.vip/e9jfjh.jpg)

```
输入：n = 4
输出：2
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：1


提示：

1 <= n <= 9
皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。

*/

/*
思路：
`深度优先搜索`


复杂度：
时间复杂度为 O(n!)
空间复杂度为 O(n)
*/

func totalNQueens(n int) int {
	res := 0
	var dfs func(n, row, col, pie, na int)

	dfs = func(n, row, col, pie, na int) {
		if row >= n {
			res++
			return
		}
		// 将所有能放置 Q 的位置由 0 变成 1，以便进行后续的位遍历
		// 也就是得到当前所有的空位
		bits := ^(col | pie | na) & ((1 << n) - 1)
		for bits > 0 {
			// 取最低位的1
			p := bits & -bits
			// 把P位置上放入皇后
			bits = bits & (bits - 1)
			// row + 1 搜索下一行可能的位置
			// col | p 目前所有放置皇后的列
			// (pie | p) << 1 和 (na | p) >> 1) 与已放置过皇后的位置 位于一条斜线上的位置
			dfs(n, row+1, col|p, (pie|p)<<1, (na|p)>>1)
		}
	}

	dfs(n, 0, 0, 0, 0)
	return res
}
