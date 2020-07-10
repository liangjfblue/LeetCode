//实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
//
// 注意：本题相对原题稍作改动
//
// 示例：
//
// 输入： 1->2->3->4->5 和 k = 2
//输出： 4
//
// 说明：
//
// 给定的 k 保证是有效的。
// Related Topics 链表 双指针

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return -1
	}

	fast, slow := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow.Val
}

//leetcode submit region end(Prohibit modification and deletion)
