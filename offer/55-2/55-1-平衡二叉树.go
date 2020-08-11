/**
 *
 * @author liangjf
 * @create on 2020/8/11
 * @version 1.0
 */
/**
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。



示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。



限制：

1 <= 树的结点个数 <= 10000
*/
package _5_2

import (
	"math"
	"offer/model"
)

var isBalancedFlag = true

func isBalanced(root *model.TreeNode) bool {
	if root == nil {
		return true
	}

	isBalancedFlag = true
	_isBalanced(root)

	return isBalancedFlag
}

func _isBalanced(root *model.TreeNode) int {
	if root == nil || !isBalancedFlag {
		return 0
	}

	l := _isBalanced(root.Left) + 1
	r := _isBalanced(root.Right) + 1
	if math.Abs(float64(l-r)) > 1 {
		isBalancedFlag = false
	}
	return int(math.Max(float64(l), float64(r)))
}
