//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
// Related Topics 链表 数学

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	h := new(ListNode)
	p := h
	p1, p2 := l1, l2
	v1, v2 := 0, 0
	for p1 != nil || p2 != nil {
		v1, v2 = 0, 0
		if p1 != nil {
			v1 = p1.Val
		}

		if p2 != nil {
			v2 = p2.Val
		}

		sum := v1 + v2 + flag
		flag = sum / 10
		sum %= 10

		p.Next = &ListNode{
			Val: sum,
		}
		p = p.Next

		if p1 != nil {
			p1 = p1.Next
		}

		if p2 != nil {
			p2 = p2.Next
		}
	}

	if flag == 1 {
		p.Next = &ListNode{
			Val: 1,
		}
	}

	return h.Next
}

//leetcode submit region end(Prohibit modification and deletion)
