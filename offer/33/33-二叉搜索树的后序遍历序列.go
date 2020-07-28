/**
 *
 * @author liangjf
 * @create on 2020/7/28
 * @version 1.0
 */
/**
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。


参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true

4, 8, 6, 12, 16, 14, 10

      10
    /   \
   6     14
  / \   / \
 4   8 12 16


提示：

数组长度 <= 1000
*/
package _3

//后序遍历的特点是根节点是最后一个元素,
//所以通过获取根节点, 把用根节点前的把前面的节点比较, 分为两半
//(因为二叉搜索树的特点是根节点的左左子树都比根节点小, 右子树都比根节点大), 因此可以以根节点切分
//递归这个过程, 就可以逐渐搜索数组, 若不符合二叉搜索树的特点就返回false
func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}

	return _verifyPostorder(postorder, 0, len(postorder)-1)
}

func _verifyPostorder(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}

	root := postorder[end]

	//以root查找左右子树分切线下标
	i := start
	for ; i < end; i++ {
		if postorder[i] > root {
			break
		}
	}

	//i的后半段是否有比root小的(正常后半段都比root大), 有就不符合
	for j := i; j < end; j++ {
		if postorder[j] < root {
			return false
		}
	}

	return _verifyPostorder(postorder, start, i-1) && _verifyPostorder(postorder, i, end-1)
}
