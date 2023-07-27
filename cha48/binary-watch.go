package main

import (
	"fmt"
	"math/bits"
)

/*
二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。

每个 LED 代表一个 0 或 1，最低位在右侧。
```

![](https://p.ipic.vip/3jkqhl.jpg)

```
例如，上面的二进制手表读取 “3:25”。

给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。

示例：

输入: n = 1
返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]

提示：

输出的顺序没有要求。
小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
超过表示范围（小时 0-11，分钟 0-59）的数据将会被舍弃，也就是说不会出现 "13:00", "0:61" 等时间。
*/

/*
思路:

复杂度：
时间复杂度：O(turnedOn^2)
空间复杂度：O(1)
*/
func readBinaryWatch(turnedOn int) (ans []string) {
	h_num := min(4, turnedOn+1)
	for i := 0; i <= h_num; i++ {
		for _, h := range possible_number(i, false) {
			m := turnedOn - i
			if m < 0 {
				continue
			}
			for _, m := range possible_number(m, true) {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return
}

func possible_number(count int, m bool) (ans []int) {
	if count == 0 {
		return []int{0}
	}
	if !m {
		for i := 0; i < 12; i++ {
			if bits.OnesCount(uint(i)) == count {
				ans = append(ans, i)
			}
		}
	} else {
		for i := 0; i < 60; i++ {
			if bits.OnesCount(uint(i)) == count {
				ans = append(ans, i)
			}
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
