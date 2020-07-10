//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
//
// 示例 1:
//
// 输入: 1->2->3->3->4->4->5
//输出: 1->2->5
//
//
// 示例 2:
//
// 输入: 1->1->1->2->3
//输出: 2->3
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
1. 以防开头就有重复元素, 所以创建哑结点, 指向头结点, 哑结点赋值给p1, 因此p1指向head
2. 利用双指针前后对比, 为了防止nil, 循环结束条件是p2 != nil && p2.Next != nil
3. 当遇到前后不一致时, 分两种情况:
	1. 当前指针p1.Next和快指针p2相等, 证明没有连续, 更新p1到快指针位置
	2. 当前指针p1.Next和快指针p2不相等, 证明连续, 直接更新p1的next为相同元素的下一个位置, 即p2的next
4. 循环结束时, 如果p1.Next != p2, 那么在最后两个元素相同, 因此特殊处理即可
*/
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var tmp = new(ListNode)
	var p1, p2 = tmp, head
	tmp.Next = head
	for p2 != nil && p2.Next != nil {
		if p2.Val != p2.Next.Val {
			if p1.Next == p2 {
				//非连续相同
				p1 = p2
			} else {
				//连续相同
				p1.Next = p2.Next
			}
		}
		p2 = p2.Next
	}

	if p1.Next != p2 {
		p1.Next = p2.Next
	}

	return tmp.Next
}

//leetcode submit region end(Prohibit modification and deletion)
