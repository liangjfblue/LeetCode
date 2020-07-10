//给定两个用链表表示的整数，每个节点包含一个数位。
// 这些数位是反向存放的，也就是个位排在链表首部。
// 编写函数对这两个整数求和，并用链表形式返回结果。
//
//
//
// 示例：
//
//
//输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912
//
//
// 进阶：假设这些数位是正向存放的，请再做一遍。
//
// 示例：
//
//
//输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
//输出：9 -> 1 -> 2，即912
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
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 != nil {
		return l2
	}

	flag, sum := 0, 0
	cur := new(ListNode)
	var head = cur
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + flag
		flag = sum / 10
		l1.Val = sum % 10
		cur.Next = l1
		cur = cur.Next
		l1, l2 = l1.Next, l2.Next
	}

	var next *ListNode
	if l1 != nil {
		next = l1
	} else {
		next = l2
	}

	for next != nil {
		sum = next.Val + flag
		flag = sum / 10
		next.Val = sum % 10
		cur.Next = next
		cur = cur.Next
		next = next.Next

	}

	if flag > 0 {
		cur.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}

	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)
