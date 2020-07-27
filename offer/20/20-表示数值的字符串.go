/**
 *
 * @author liangjf
 * @create on 2020/7/27
 * @version 1.0
 */
/**
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"0123"都表示数值，
但"12e"、"1a3.14"、"1.2.3"、"+-5"、"-1E-16"及"12e+5.4"都不是。
*/
package _0

import "strings"

//不符合的类型:
//1.有非e的字符
//2.有e, 但e后面没有数字
//3.有e, e后面有非数字字符
//4.开头有多个正/负字符
//5.有多个小数点
//时间复杂度O(n), 空间复杂度O(1)
func isNumber(s string) bool {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return false
	}

	if len(s) == 1 && s[0] < '0' || s[0] > '9' {
		return false
	}

	eFlag := false
	numFlag := false
	dotFlag := false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			//数字
			numFlag = true
		} else if s[i] == '.' && !dotFlag && !eFlag {
			//当前是., 没有出现过.并且没有出现过e
			dotFlag = true
		} else if (s[i] == 'e' || s[i] == 'E') && numFlag && !eFlag {
			//当前是e, 已出现过数字并且首次出现e
			eFlag = true
			numFlag = false //出现过e标记为false, 为了判断123e, e后面必须有数字
		} else if (s[i] == '-' || s[i] == '+') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E') {
			//开头是+-
		} else {
			return false
		}
	}

	return numFlag
}
