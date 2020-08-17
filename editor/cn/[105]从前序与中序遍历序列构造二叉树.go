//根据一棵树的前序遍历与中序遍历构造二叉树。
//
// 注意:
//你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//
// 返回如下的二叉树：
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// Related Topics 树 深度优先搜索 数组
// 👍 622 👎 0

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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}

	indexM := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		indexM[inorder[i]] = i
	}

	//	return _buildTree1(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1, indexM)
	return _buildTree2(preorder, inorder)
}

func _buildTree1(
	preorder, inorder []int,
	preorderLeft, preorderRight, inorderLeft, inorderRight int,
	indexM map[int]int,
) *TreeNode {
	if preorderLeft > preorderRight {
		return nil
	}

	//当前子树的根节点
	root := preorder[preorderLeft]

	inorderRoot := indexM[root]

	//左子树的节点个数是
	leftTreeNodes := inorderRoot - inorderLeft

	//构造子树的根节点
	node := &TreeNode{Val: root}

	//前序根据中序的根前后的左右节点个数来切分, 中序根据根节点所在的位置来切分左右子树
	//递归构建左子树
	node.Left = _buildTree1(
		preorder,
		inorder,
		preorderLeft+1,
		preorderLeft+leftTreeNodes,
		inorderLeft,
		inorderRoot-1,
		indexM,
	)
	//递归构建右子树
	node.Right = _buildTree1(
		preorder,
		inorder,
		preorderLeft+leftTreeNodes+1,
		preorderRight,
		inorderRoot+1,
		inorderRight,
		indexM,
	)

	return node
}

func _buildTree2(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	// 前序遍历 preorder =[3,9,20,15,7]
	//中序遍历 inorder   =[9,3,15,20,7]
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
		if inorder[inorderRoot] == preorder[0] {
			break
		}
	}

	root.Left = _buildTree2(preorder[1:inorderRoot+1], inorder[:inorderRoot])
	root.Right = _buildTree2(preorder[inorderRoot+1:], inorder[inorderRoot+1:])

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
