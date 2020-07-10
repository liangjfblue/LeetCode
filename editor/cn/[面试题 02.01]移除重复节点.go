//编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
//
// 示例1:
//
//
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
//
//
// 示例2:
//
//
// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
//
//
// 提示：
//
//
// 链表长度在[0, 20000]范围内。
// 链表元素在[0, 20000]范围内。
//
//
// 进阶：
//
// 如果不得使用临时缓冲区，该怎么解决？
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
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		nodeM = make(map[int]bool)
		cur   = head
	)

	for cur != nil && cur.Next != nil {
		nodeM[cur.Val] = true
		if _, ok := nodeM[cur.Next.Val]; ok {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
