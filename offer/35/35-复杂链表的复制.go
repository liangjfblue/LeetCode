/**
 *
 * @author liangjf
 * @create on 2020/7/29
 * @version 1.0
 */
/**
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
还有一个 random 指针指向链表中的任意节点或者 null。

输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
*/
package _5

import "offer/model"

//哈希法:
//	遍历链表, 用hash记录next和random指针的关系
//空间和时间都是O(n)
func copyRandomList(head *model.Node) *model.Node {
	if head == nil {
		return head
	}

	//map存(原节点，拷贝节点)的一个映射
	cur := head
	nodeM := make(map[*model.Node]*model.Node)
	for cur != nil {
		nodeM[cur] = &model.Node{Val: cur.Val}
		cur = cur.Next
	}

	//将拷贝的新的节点组织成一个链表(借助cur的Next,Random指针)
	cur = head
	for cur != nil {
		nodeM[cur].Next = nodeM[cur.Next]
		nodeM[cur].Random = nodeM[cur.Random]
		cur = cur.Next
	}

	return nodeM[head]
}

//原地法:
//	1.复制原链表, 各节点放在被复制节点的后面
//  2.复制节点设好Random指针关系
//	3.分离链表
//主要是利用了跨节点来访问原链表和clone链表
func copyRandomList2(head *model.Node) *model.Node {
	if head == nil {
		return head
	}

	//复制node放到原链表对应节点后一位
	for cur := head; cur != nil; cur = cur.Next.Next {
		node := &model.Node{Val: cur.Val}
		node.Next = cur.Next
		cur.Next = node
	}

	//复制Random
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			//clone节点  原节点
			cur.Next.Random = cur.Random.Next
		}
	}

	//分离链表
	var next *model.Node
	newHead := head.Next
	for cur := head; cur != nil && cur.Next != nil; {
		next = cur.Next
		cur.Next = cur.Next.Next
		cur = next
	}

	return newHead
}
