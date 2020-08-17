//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
// 示例:
//			   [-10,-3,0,5,9,10]
// 给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
//
// Related Topics 树 深度优先搜索
// 👍 548 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return build(nums)
}

func build(nums []int) *TreeNode {
	sum := len(nums)
	if sum == 0 {
		return nil
	}

	//找出mid节点
	mid := sum / 2

	root := new(TreeNode)
	root.Val = nums[mid]

	//递归构建左子树
	root.Left = build(nums[:mid])
	//递归构建右子树
	root.Right = build(nums[mid+1:])
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
