/**
 *
 * @author liangjf
 * @create on 2020/7/27
 * @version 1.0
 */
/**
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000
*/
package _5

import "offer/model"

func mergeTwoLists(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	newHead := new(model.ListNode)
	cur := newHead
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			cur.Next = p1
			cur = cur.Next
			p1 = p1.Next
		} else {
			cur.Next = p2
			cur = cur.Next
			p2 = p2.Next
		}
	}

	if p1 != nil {
		cur.Next = p1
	}

	if p2 != nil {
		cur.Next = p2
	}

	return newHead.Next
}
