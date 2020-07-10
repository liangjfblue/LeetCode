//给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
//
// 示例 1:
//
// 输入: 1->2->3->4->5->NULL, k = 2
//输出: 4->5->1->2->3->NULL
//解释:
//向右旋转 1 步: 5->1->2->3->4->NULL
//向右旋转 2 步: 4->5->1->2->3->NULL
//
//
// 示例 2:
//
// 输入: 0->1->2->NULL, k = 4
//输出: 2->0->1->NULL
//解释:
//向右旋转 1 步: 2->0->1->NULL
//向右旋转 2 步: 1->2->0->NULL
//向右旋转 3 步: 0->1->2->NULL
//向右旋转 4 步: 2->0->1->NULL
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
/**
思路:
1.获取链表长度n, 计算出需要移动的步长 k%n
2.快慢指针, 快指针先走k%n
3.然后快慢指针一起走, 直至快指针为nil, 此时慢指针指向的后半段就是需要挪到前面的, 改变指针指向即可
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	var fast, slow = head, head

	sum := 0
	for fast != nil {
		sum++
		fast = fast.Next
	}

	//真正移动的距离
	k %= sum

	//快指针先走k步
	fast = head
	for k > 0 {
		k--
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	fast.Next = head
	head = slow.Next
	slow.Next = nil

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
