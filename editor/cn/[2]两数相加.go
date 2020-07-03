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
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	}

	carry := 0
	head := new(ListNode)
	var cur = head

	//只要任意链表不为空, 或者有进位就继续
	for l1 != nil || l2 != nil || carry == 1 {
		if l1 != nil {
			cur.Val += l1.Val
		}
		if l2 != nil {
			cur.Val += l2.Val
		}
		cur.Val += carry

		carry = cur.Val / 10
		cur.Val = cur.Val % 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		//防止最后一次new
		if l1 != nil || l2 != nil || carry == 1 {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
