//根据一棵树的中序遍历与后序遍历构造二叉树。
//
// 注意:
//你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 中序遍历 inorder =  [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]
//
// 返回如下的二叉树：
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// Related Topics 树 深度优先搜索 数组
// 👍 264 👎 0

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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || postorder == nil {
		return nil
	}
	return _buildTree(inorder, postorder)
}

func _buildTree(inorder, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	//后序遍历最后一个元素是根节点
	root := &TreeNode{Val: postorder[len(postorder)-1]}

	//中序遍历 inorder =  [9,3,15,20,7]
	//后序遍历 postorder= [9,15,7,20,3]
	//
	// 返回如下的二叉树：
	//
	//     3
	//   / \
	//  9  20
	//    /  \
	//   15   7
	inorderRoot := 0
	for ; inorderRoot < len(inorder); inorderRoot++ {
		if inorder[inorderRoot] == postorder[len(postorder)-1] {
			break
		}
	}

	root.Left = _buildTree(inorder[:inorderRoot], postorder[:inorderRoot])
	root.Right = _buildTree(inorder[inorderRoot+1:], postorder[inorderRoot:len(postorder)-1])

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
