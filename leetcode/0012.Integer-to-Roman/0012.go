package leetcode

import "fmt"
func intToRoman(num int) string {
    // 記錄所有羅馬字母對應的數字, 包含以下要額外紀錄
    // I 可以放在 V (5) 和 X (10) 的左邊，來表示 4 和 9。
    // X 可以放在 L (50) 和 C (100) 的左邊，來表示 40 和 90。
    // C 可以放在 D (500) 和 M (1000) 的左邊，來表示 400 和 900。
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
        // 不斷 i++, values[i] < num 才跳出迴圈, 如果是3, values[i] = 1, 就會跳出迴圈
		for values[i] > num {
            fmt.Println("values[i]: ", values[i])
			i++
		}
        // fmt.Println("symbols[i]: ", symbols[i])
        // num - 已經計算的羅馬字母的數字
		num -= values[i]
        // res 已經計算的羅馬字母
		res += symbols[i]
	}
	return res
}