/**
 *
 * @author liangjf
 * @create on 2020/7/27
 * @version 1.0
 */
/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL


限制：

0 <= 节点个数 <= 5000
*/
package _4

import "offer/model"

//迭代法:
//pre会成为新的头结点
//迭代操作:
//	保存好next指针
//	后一个节点指向前一个节点
//	移动前后节点
//	循环...
func reverseList(head *model.ListNode) *model.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	var pre, next *model.ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

//递归法: 第二节点是反转后的头节点
//递归操作:
//	反转前后节点
func reverseList2(head *model.ListNode) *model.ListNode {
	if head == nil {
		return head
	}

	node := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node
}
