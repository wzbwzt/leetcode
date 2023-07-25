package main

import "math"

/*
给你一个字符串 S、一个字符串 T 。请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。



示例：

输入：S = "ADOBECODEBANC", T = "ABC"
输出："BANC"


提示：

如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/

/*
思路：
`滑动窗口`
while 右边界 < 合法条件：
  # 右边界扩张
  window右边界+1
  更新状态信息
  # 左边界收缩
  while 符合收缩条件：
    window左边界+1
    更新状态信息

复杂度：
时间复杂度： O(N+K);N 为 S 串长度，K 为 T 串长度
空间复杂度： O(S)，其中 S 为 T 字符集元素个数
*/

func minWindow(s string, t string) (ans string) {
	counter := map[byte]int{}
	l := 0
	N := len(s)
	counter_t := map[byte]int{}
	for i := 0; i < len(t); i++ {
		counter_t[t[i]]++
	}

	res := math.MaxInt32
	k := 0
	for r := 0; r < N; r++ {
		counter[s[r]]++
		if _, ok := counter_t[s[r]]; ok && counter[s[r]] == counter_t[s[r]] {
			k++
		}
		for k == len(counter_t) {
			//收缩左边
			if r-l+1 < res {
				ans = s[l : r+1]
			}
			res = min((r - l + 1), res)

			counter[s[l]]--
			if _, ok := counter_t[s[l]]; ok && counter[s[l]] == counter_t[s[l]]-1 {
				k--
			}
			l++
		}

	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
