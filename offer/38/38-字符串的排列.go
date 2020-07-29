/**
 *
 * @author liangjf
 * @create on 2020/7/29
 * @version 1.0
 */
/**
输入一个字符串，打印出该字符串中字符的所有排列。


你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。


示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

限制：
1 <= s 的长度 <= 8
*/
package _8

func permutation(s string) []string {
	if len(s) == 0 {
		return nil
	}

	res := make([]string, 0)
	bfs(&res, []byte(s), 0)

	return res
}

func bfs(res *[]string, tmp []byte, i int) {
	if i == len(tmp)-1 {
		*res = append(*res, string(tmp))
	}

	dic := make(map[byte]bool)
	for j := i; j < len(tmp); j++ {
		//剪枝
		if _, ok := dic[tmp[j]]; ok {
			continue
		}

		//加入已访问dic(整个递归过程都共享)
		dic[tmp[j]] = true
		tmp[i], tmp[j] = tmp[j], tmp[i] // 交换，将 c[i] 固定在第 x 位
		bfs(res, tmp, i+1)              //开启固定第 x + 1 位字符
		tmp[i], tmp[j] = tmp[j], tmp[i] // 恢复交换
	}
}
