/**
 *
 * @author liangjf
 * @create on 2020/7/27
 * @version 1.0
 */
/**
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。



示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.


提示：

各函数的调用总次数不超过 20000 次
*/
package _0

//思路:
//1.和普通的栈一样, 但是元素是一个结构体, 有一个min字段记录当前栈的最小值
//2.获取min时, 直接就是栈顶的元素的min
type data struct {
	d   int
	min int
}

type MinStack struct {
	data []data
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{data: make([]data, 0)}
}

func (this *MinStack) Push(x int) {
	if len(this.data) == 0 {
		this.data = append(this.data, data{
			d:   x,
			min: x,
		})
		return
	}

	min := this.data[len(this.data)-1].min
	if x < min {
		min = x
	}

	this.data = append(this.data, data{
		d:   x,
		min: min,
	})
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}

	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1].d
}

func (this *MinStack) Min() int {
	return this.data[len(this.data)-1].min
}
