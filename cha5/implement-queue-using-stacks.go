package main

/*
使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。
示例:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek(); // 返回 1
queue.pop(); // 返回 1
queue.empty(); // 返回 false
说明:

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的、 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
*/

/*
思路：
栈：先入后出
队列：先入先出
使用双栈的来解决:instack,outstack
instack:每次push时入栈
outstack:每次pop/peek时，弹栈；如果为空，先从instack依次弹栈，再入栈outstack

复杂度：
空间复杂度：O(n)
时间复杂度：O(n)
*/

type MyQueue struct {
	instack, outstack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (m *MyQueue) Push(x int) {
	m.instack = append(m.instack, x)
}

func (m *MyQueue) in2out() {
	for i := len(m.instack) - 1; i >= 0; i-- {
		m.outstack = append(m.outstack, m.instack[i])
	}
	m.instack = nil
}
func (m *MyQueue) Pop() int {
	if len(m.outstack) == 0 {
		m.in2out()
	}
	out := m.outstack[len(m.outstack)-1]
	m.outstack = m.outstack[:len(m.outstack)-1]
	return out
}

func (m *MyQueue) Peek() int {
	if len(m.outstack) == 0 {
		m.in2out()
	}
	return m.outstack[len(m.outstack)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.instack) == 0 && len(m.outstack) == 0
}
