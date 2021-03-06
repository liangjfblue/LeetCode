//给定一个有环链表，实现一个算法返回环路的开头节点。 有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。 示例 1
//： 输入：head = [3,2,0,-4], pos = 1 输出：tail connects to node index 1 解释：链表中有一个环，其尾部连
//接到第二个节点。 示例 2： 输入：head = [1,2], pos = 0 输出：tail connects to node index 0 解释：链表中有
//一个环，其尾部连接到第一个节点。 示例 3： 输入：head = [1], pos = -1 输出：no cycle 解释：链表中没有环。 进阶： 你是否可以不
//用额外空间解决此题？ Related Topics 链表

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast != slow {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

//leetcode submit region end(Prohibit modification and deletion)
