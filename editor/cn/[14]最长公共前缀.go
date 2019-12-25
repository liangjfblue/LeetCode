//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1:
//
// 输入: ["flower","flow","flight"]
//输出: "fl"
//
//
// 示例 2:
//
// 输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//
// 说明:
//
// 所有输入只包含小写字母 a-z 。
// Related Topics 字符串
package cn

import (
	"sync"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
思路：
1、若输入为0个子串，直接返回""
2、若输入为1和子串，直接返回子串
3、其他情况：
	1、拿出strs[0],作为比较项；
	2、遍历除子串1的字符串列表，并创建n个goroutine，传入strs[0]，和各自的字符串str[i]
	3、利用sync.WaitGroup和chan，等待goroutine完成和遍历chan得到所有子串中最小的公共前缀下标
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var (
		masterStr = strs[0]
		ch        = make(chan int, len(strs)-1)
		wait      sync.WaitGroup
		min       = 0xfffffff
	)

	for _, str := range strs[1:] {
		wait.Add(1)
		go func(s1, s2 string, c chan int) {
			defer wait.Done()
			index := 0
			for i := 0; i < len(s1) && i < len(s2); i++ {
				index++
				if s1[i] != s2[i] {
					c <- i
					return
				}
			}
			c <- index
		}(masterStr, str, ch)
	}

	wait.Wait()
	close(ch)

	for c := range ch {
		if c < min {
			min = c
		}
	}

	return masterStr[:min]
}

//leetcode submit region end(Prohibit modification and deletion)
