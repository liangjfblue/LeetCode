/**
 *
 * @author liangjf
 * @create on 2020/7/24
 * @version 1.0
 */
/**
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

限制：

0 <= 节点个数 <= 5000
*/
package _7

import "offer/model"

//前序遍历的第一个元素肯定是子树的根, 根据前序的根在中序的位置分为左右子树两半, 并且从中序中得到左右子树的个数
//重复这个过程就可以还原一个二叉树
//时间复杂度O(n), 空间复杂度O(n)
func buildTree(preorder []int, inorder []int) *model.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	//前序第一个元素是子树的根
	tree := &model.TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	//从中序找出根的左右子树分界线下标, 得到左右子树的个数, 从而去前序计算出左右子树的区间(为了递归得到左右子树的第一个元素是根)
	idx := 0
	for k, v := range inorder {
		if v == tree.Val {
			idx = k
		}
	}

	//左右子树递归
	//(前序: 左子树递归区间:下一个元素~前序首元素所在中序的位置的下标; 右子树递归区间:前序首元素所在中序的位置的下标+1~末尾)
	//(中序: 直接根据前序首元素所在中序的位置的下标分半)
	tree.Left = buildTree(preorder[1:idx+1], inorder[0:idx])
	tree.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return tree
}
