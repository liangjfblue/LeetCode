//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
// 说明:
//1 ≤ m ≤ n ≤ 链表长度。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL
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
/**
思路:
1. 遍历到反转前, 保存节点信息
2. 反转区间节点
3. 用之前保存的节点来连接反转后的区间
*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	var count = 1
	var pre0 = head
	var cur0 = head
	var pre *ListNode
	var cur = head

	//循环遍历 反转区间 直至nil
	for cur != nil && count <= n {
		//反转前的节点信息
		if count == m {
			pre0 = pre
			cur0 = cur
		}

		//需要反转的区间
		if count > m && count <= n {
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		} else {
			//反转前的子链表迭代遍历即可
			pre = cur
			cur = cur.Next
		}
		count++
	}
	if pre0 == nil {
		head = pre
	} else {
		pre0.Next = pre
	}

	cur0.Next = cur

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
