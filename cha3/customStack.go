package main

/*
请你设计一个支持下述操作的栈。

实现自定义栈类 CustomStack ：

CustomStack(int maxSize)：用 maxSize 初始化对象，maxSize 是栈中最多能容纳的元素数量，栈在增长到 maxSize 之后则不支持 push 操作。
void push(int x)：如果栈还未增长到 maxSize ，就将 x 添加到栈顶。
int pop()：弹出栈顶元素，并返回栈顶的值，或栈为空时返回 -1 。
void inc(int k, int val)：栈底的 k 个元素的值都增加 val 。如果栈中元素总数小于 k ，则栈中的所有元素都增加 val 。


示例：

输入：
["CustomStack","push","push","pop","push","push","push","increment","increment","pop","pop","pop","pop"]
[[3],[1],[2],[],[2],[3],[4],[5,100],[2,100],[],[],[],[]]
输出：
[null,null,null,2,null,null,null,null,null,103,202,201,-1]
解释：
CustomStack customStack = new CustomStack(3); // 栈是空的 []
customStack.push(1); // 栈变为 [1]
customStack.push(2); // 栈变为 [1, 2]
customStack.pop(); // 返回 2 --> 返回栈顶值 2，栈变为 [1]
customStack.push(2); // 栈变为 [1, 2]
customStack.push(3); // 栈变为 [1, 2, 3]
customStack.push(4); // 栈仍然是 [1, 2, 3]，不能添加其他元素使栈大小变为 4
customStack.increment(5, 100); // 栈变为 [101, 102, 103]
customStack.increment(2, 100); // 栈变为 [201, 202, 103]
customStack.pop(); // 返回 103 --> 返回栈顶值 103，栈变为 [201, 202]
customStack.pop(); // 返回 202 --> 返回栈顶值 202，栈变为 [201]
customStack.pop(); // 返回 201 --> 返回栈顶值 201，栈变为 []
customStack.pop(); // 返回 -1 --> 栈为空，返回 -1


提示：

1 <= maxSize <= 1000
1 <= x <= 1000
1 <= k <= 1000
0 <= val <= 100
每种方法
*/

/*
思路:
前缀和思路：维护一个increDiff固定容量的数组，记录每次incre 操作的差值
push时 increDiff: append(0)
increment(k,v)时 increDiff[k-1]+=v  (k =min(k,cnt))
pop时 out=list[cnt-1]+incretment[cnt-1]
increment[cnt-2]+=increment[cnt-1]
re切割list,increment


时间复杂度：全部都是 O(1)
空间复杂度为 O(cnt)
*/

type CustomStack struct {
	list      []int
	cnt       int
	maxsize   int
	increDiff []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{list: make([]int, 0, maxSize), cnt: 0, maxsize: maxSize, increDiff: make([]int, 0, maxSize)}
}

func (c *CustomStack) Push(x int) {
	if c.cnt == c.maxsize {
		return
	}
	c.list = append(c.list, x)
	c.increDiff = append(c.increDiff, 0)
	c.cnt++
}

func (c *CustomStack) Pop() int {
	if c.cnt == 0 {
		return -1
	}
	out := c.list[c.cnt-1] + c.increDiff[c.cnt-1]

	c.list = c.list[:c.cnt-1]
	if c.cnt >= 2 {
		c.increDiff[c.cnt-2] += c.increDiff[c.cnt-1]
	}
	c.increDiff = c.increDiff[:c.cnt-1]

	c.cnt--
	return out
}

func (c *CustomStack) Increment(k int, val int) {
	if k > c.cnt {
		k = c.cnt
	}
	if k == 0 {
		return
	}
	c.increDiff[k-1] += val
}
