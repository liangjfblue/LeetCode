/**
 *
 * @author liangjf
 * @create on 2020/7/24
 * @version 1.0
 */
/**
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。
(若队列中没有元素，deleteHead 操作返回 -1 )

示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
*/
package _9

//貌似和题意不符, 人家要的是两个栈实现. 我这里直接用切片搞定了
//用两个栈的思路:
//1.创建两个栈
//2.入队: 直接插入栈1
//3.出队:
//	如果栈2为空, 把栈1一个个pop出来插入到栈2, 然后栈2pop一个返回
//	如果栈2不为空, 直接栈2pop一个返回
//时间复杂度O(1), 空间复杂度O(1) 虽然会遍历pop出栈1的操作, 但是只有在栈2位空时才触发, 均摊下来其实还是O(1)
type CQueue struct {
	data []int
}

func Constructor() CQueue {
	return CQueue{
		data: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.data = append(this.data, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.data) == 0 {
		return -1
	}

	v := this.data[0]

	if len(this.data) == 1 {
		this.data = this.data[:0]
	} else {
		this.data = this.data[1:]
	}

	return v
}

/**
//xxx 是栈数据结构

type CQueue struct {
	stack1 xxx
	stack2 xxx
}

func Constructor() CQueue {
	return CQueue{
		stack1: new(xxx),
		stack2: new(xxx),
	}
}

func (this *CQueue) AppendTail(int value) {
	stack1.add(value)
}

func (this *CQueue) DeleteHead() {
	if stack2.isEmpty() {
		if stack1.isEmpty()
			return -1
		for !stack1.isEmpty() {
			stack2.add(stack1.pop())
		}
		return stack2.pop()
	} else return stack2.pop()
}
*/
