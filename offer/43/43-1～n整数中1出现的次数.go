/**
 *
 * @author liangjf
 * @create on 2020/7/30
 * @version 1.0
 */
/**
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。


示例 1：

输入：n = 12
输出：5
示例 2：

输入：n = 13
输出：6


限制：

1 <= n < 2^31
*/
package _3

func countDigitOne(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		tmp := i
		for tmp != 0 {
			if tmp%10 == 1 {
				sum++
			}
			tmp /= 10
		}
	}
	return sum
}
