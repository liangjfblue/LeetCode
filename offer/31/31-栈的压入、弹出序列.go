/**
 *
 * @author liangjf
 * @create on 2020/7/27
 * @version 1.0
 */
/**
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。
例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。


示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：

输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。


提示：

0 <= pushed.length == popped.length <= 1000
0 <= pushed[i], popped[i] < 1000
pushed 是 popped 的排列。
*/
package _1

//加多一个辅助栈, 把pushed压栈,
//一边压栈一边和popped当前元素对比,
//如果相等那么就pop出pushed栈顶元素, 并且popped后移一位,
//满足条件是最终位移和popped长度一致
func validateStackSequences(pushed []int, popped []int) bool {
	popIndex := 0
	popLen := len(popped)
	stack := make([]int, 0)

	for _, v := range pushed {
		stack = append(stack, v)
		for popIndex < popLen && len(stack) > 0 && stack[len(stack)-1] == popped[popIndex] {
			stack = stack[:len(stack)-1]
			popIndex++
		}
	}

	return popIndex == popLen
}
