//给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
// 示例:
//
// 给定的有序链表： [-10, -3, 0, 5, 9],
//
//一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
//
// Related Topics 深度优先搜索 链表

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
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
1.找出链表中点
2.递归以链表中心为分割的前后子链表, 构造二叉树
*/
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	//找出链表中点 123456 12345
	var pre *ListNode
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}

	node := &TreeNode{
		Val:   slow.Val,
		Left:  nil,
		Right: nil,
	}

	if head == slow {
		return node
	}

	if pre != nil {
		pre.Next = nil
	}

	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(slow.Next)
	return node
}

//leetcode submit region end(Prohibit modification and deletion)
