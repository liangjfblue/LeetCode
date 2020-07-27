/**
 *
 * @author liangjf
 * @create on 2020/7/27
 * @version 1.0
 */
/**
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。


示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.
*/
package _2

import "offer/model"

func getKthFromEnd(head *model.ListNode, k int) *model.ListNode {
	if head == nil || (head.Next == nil && k == 1) {
		return head
	}

	//先移动k格
	p := head
	for p != nil && k > 0 {
		p = p.Next
		k--
	}

	//若已经超出, 返回nil
	if p == nil {
		return nil
	}

	//截取前面n-k个
	cur := head
	for p != nil {
		cur = cur.Next
		p = p.Next
	}

	head = cur

	return head
}
