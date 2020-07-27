/**
 *
 * @author liangjf
 * @create on 2020/7/24
 * @version 1.0
 */
/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."

限制：

0 <= s 的长度 <= 10000
*/
package _5

//遍历字符串, 追加到新的字符串中, 遇到空格就替换而已
//时间复杂度O(n), 空间复杂度O(n)
func replaceSpace(s string) string {
	if s == "" {
		return ""
	}

	ret := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ret = append(ret, "%20"...)
		} else {
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}

//时间复杂度O(n), 空间复杂度O(1)
func replaceSpace2(s string) string {
	if s == "" {
		return ""
	}

	for i := 0; i < len(s); {
		if s[i] == ' ' {
			s = s[:i] + "%20" + s[i+1:]
			i += 3
		} else {
			i++
		}
	}
	return s
}
