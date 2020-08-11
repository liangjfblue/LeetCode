/**
 *
 * @author liangjf
 * @create on 2020/8/10
 * @version 1.0
 */
/**
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。



示例 1:

输入: [7,5,6,4]
输出: 5


限制：

0 <= 数组长度 <= 50000
*/
package _1

//暴力法
func reversePairs2(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				sum++
			}
		}
	}

	return sum
}

//归并法
func reversePairs(nums []int) int {
	sum := 0
	mergeSort(&nums, 0, len(nums)-1, &sum)
	return sum
}

func mergeSort(nums *[]int, left, right int, sum *int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2

	mergeSort(nums, left, mid, sum)
	mergeSort(nums, mid+1, right, sum)
	merge(nums, left, mid, right, sum)
	return
}

func merge(nums *[]int, left, mid, right int, sum *int) {
	i, j := left, mid+1
	tmp := make([]int, 0, right-left)
	for i <= mid && j <= right {
		if (*nums)[i] > (*nums)[j] {
			//规定排序, 左边有序, 右边有序, 所以左边当i下标已经大于右边j下标时, 左边剩下的都比右边的大
			//因此, 把各个位的逆序对加起来, 就是总的逆序对
			*sum += mid - i + 1
		}

		if (*nums)[i] > (*nums)[j] {
			tmp = append(tmp, (*nums)[j])
			j++
		} else {
			tmp = append(tmp, (*nums)[i])
			i++
		}
	}

	for i <= mid {
		tmp = append(tmp, (*nums)[i])
		i++
	}

	for j <= right {
		tmp = append(tmp, (*nums)[j])
		j++
	}

	for i := 0; i < len(tmp); i++ {
		(*nums)[left+i] = tmp[i]
	}
	tmp = nil

	return
}
