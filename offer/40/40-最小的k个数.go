/**
 *
 * @author liangjf
 * @create on 2020/7/29
 * @version 1.0
 */
/**
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。


示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]


限制：

0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000
*/
package _0

import (
	"sort"
)

//排序法
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)

	return arr[:k]
}

//快排, 找出第k大, 前面的就是前k个最小的
func getLeastNumbers2(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}

	return quickSort(arr, 0, len(arr)-1, k-1)
}

//快排就是以基准值为比较, 递归的分为左右两半, 直至满足结束条件.
//过程就是挖洞填洞
func quickSort(arr []int, l, r, k int) []int {
	idx := partition(arr, l, r)
	if l >= r || idx == k {
		return arr[:idx+1]
	}

	//根据下标idx与k的大小关系来决定继续切分左段还是右段
	if idx > k {
		return quickSort(arr, l, idx-1, k)
	} else {
		return quickSort(arr, idx+1, r, k)
	}
}

func partition(arr []int, l, r int) int {
	v := arr[l]
	for l < r {
		//找出右边首个可以填入左边的洞的元素(因为有左边你首元素作为基准值, 空出一个位置)
		for arr[r] >= v && r > l {
			r--
		}
		arr[l] = arr[r]

		//找出左边首个可以填入右边的洞的元素(因为上面填洞, 挖洞后右边空出一个位置)
		for arr[l] <= v && r > l {
			l++
		}
		arr[r] = arr[l]
	}

	//交换基准值和中点
	arr[l] = v

	return l
}
