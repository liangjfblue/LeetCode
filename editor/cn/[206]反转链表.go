//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表
// 👍 1163 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p := head
	var pre *ListNode
	for p != nil {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}

	return pre
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return p
}

//leetcode submit region end(Prohibit modification and deletion)
