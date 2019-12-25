//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串
package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
思路：
1、临时map，存放左括号
2、遍历字符串，匹配最顶层同类型的左括号，是就删除切片中的左括号，否就返回false
3、最后判断切片是否有元素，有表示有未匹配的左括号，返回false
*/
func isValid(s string) bool {
	if s == "" {
		return true
	}

	sum := len(s)
	tmp := map[int32]int32{'(': ')', '{': '}', '[': ']'}
	stack := make([]uint8, 0, sum)

	for i := 0; i < sum; i++ {
		if _, ok := tmp[int32(s[i])]; ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			if tmp[int32(stack[len(stack)-1])] == int32(s[i]) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	for _, v := range stack {
		if v != 0 {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
