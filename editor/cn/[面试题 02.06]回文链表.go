//编写一个函数，检查输入的链表是否是回文的。
//
//
//
// 示例 1：
//
// 输入： 1->2
//输出： false
//
//
// 示例 2：
//
// 输入： 1->2->2->1
//输出： true
//
//
//
//
// 进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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
1.快慢指针, 找到链表终点
2.反转后半链表
3.分别遍历比较前后半段链表是否一致
4.根据是否需要还原链表
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	isPalindrome := true

	//链表长度
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	//将前半部分反转
	step := length / 2
	var prev *ListNode
	cur := head
	for i := 1; i <= step; i++ {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	mid := cur

	var left, right *ListNode = prev, nil
	if length%2 == 0 {
		//长度为偶数
		right = mid
	} else {
		right = mid.Next
	}

	//从中间向左右两边遍历对比
	for left != nil && right != nil {
		if left.Val != right.Val {
			//值不相等,不是回文链表
			isPalindrome = false
			break
		}
		left = left.Next
		right = right.Next
	}

	//将前半部分反转的链表进行复原
	cur = prev
	prev = mid
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	return isPalindrome
}

//leetcode submit region end(Prohibit modification and deletion)
