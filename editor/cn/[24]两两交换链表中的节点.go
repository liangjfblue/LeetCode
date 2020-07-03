//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例:
//
// 给定 1->2->3->4->5->6, 你应该返回 2->1->4->3->6->5.
//
// Related Topics 链表

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//递归
	/**
	思路:递归返回是返回交换好的子链表的头结点, 递归结束条件是最后一对交换, 递归操作是交换节点(前节点指向交换好的子链表, 后节点指向前节点)
	*/
	//var p2 = head.Next
	//
	//head.Next = swapPairs(p2.Next)
	//p2.Next = head
	//
	//return p2

	//迭代
	/**
	思路:三指针法, cur一直指向即将交换的前一个节点, p1,p2指向交换的两个节点, 一直双双交换直至最后一对交换结束
	*/
	var newHead = head.Next
	var cur, p1, p2 = head, head, head.Next
	for head != nil && head.Next != nil {
		p1, p2 = head, head.Next

		cur.Next = p2
		p1.Next = p2.Next
		p2.Next = p1

		cur = p1
		head = p1.Next
	}
	return newHead
}

//leetcode submit region end(Prohibit modification and deletion)
