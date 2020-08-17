//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//
//
// 示例：
//
// 给你这个链表：1->2->3->4->5
//
// 当 k = 2 时，应当返回: 2->1->4->3->5
//
// 当 k = 3 时，应当返回: 3->2->1->4->5
//
//
//
// 说明：
//
//
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
// Related Topics 链表
// 👍 684 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	p := head

	//找到本轮翻转的最后一个节点
	for p != nil && count < k {
		count++
		p = p.Next
	}

	//不足k个直接返回
	if count < k {
		return head
	}

	pp := head
	var pre *ListNode
	for pp != p {
		next := pp.Next
		pp.Next = pre
		pre = pp
		pp = next
	}

	head.Next = reverseKGroup(p, k)

	//新链表的第2个是头节点
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)
