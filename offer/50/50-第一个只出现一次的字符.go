/**
 *
 * @author liangjf
 * @create on 2020/7/31
 * @version 1.0
 */
/**
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例:

s = "abaccdeff"
返回 "b"

s = ""
返回 " "


限制：

0 <= s 的长度 <= 50000
*/
package _0

func firstUniqChar(s string) byte {
	if s == "" {
		return byte(' ')
	}

	c := make([]int, 26)

	for i := 0; i < len(s); i++ {
		c[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if c[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return byte(' ')
}
