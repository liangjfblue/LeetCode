//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
// 示例 1：
//
// 输入：head = [1,3,2]
//输出：[2,3,1]
//
//
//
// 限制：
//
// 0 <= 链表长度 <= 10000
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
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return []int{head.Val}
	}

	ret := make([]int, 0)
	cur := head
	var pre *ListNode
	//1->2->3->4 4->3->2->1
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	for pre != nil {
		ret = append(ret, pre.Val)
		pre = pre.Next
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
