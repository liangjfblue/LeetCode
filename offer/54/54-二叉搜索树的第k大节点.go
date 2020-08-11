/**
 *
 * @author liangjf
 * @create on 2020/8/11
 * @version 1.0
 */
package _4

import "offer/model"

//中序遍历
func kthLargest(root *model.TreeNode, k int) int {
	if root == nil {
		return 0
	}

	res := make([]int, 0)
	find(root, &res)
	return res[len(res)-k]
}

func find(root *model.TreeNode, res *[]int) {
	if root == nil {
		return
	}

	find(root.Left, res)
	*res = append(*res, root.Val)
	find(root.Right, res)
	return
}

func kthLargest2(root *model.TreeNode, k int) int {
	if root == nil {
		return 0
	}

	count, result := 0, 0
	find2(root, k, &count, &result)
	return result
}

func find2(root *model.TreeNode, k int, count, result *int) {
	if root == nil {
		return
	}

	find2(root.Right, k, count, result)
	*count++
	if k == *count {
		*result = root.Val
	}
	find2(root.Left, k, count, result)
}
