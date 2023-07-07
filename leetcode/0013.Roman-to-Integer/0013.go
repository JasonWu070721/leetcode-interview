package leetcode
func romanToInt(s string) int {
    // 羅馬數字包含以下七種字元: I， V， X， L，C，D 和 M。
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
        // char 從最右邊執行到最左邊
		char := s[len(s)-(i+1) : len(s)-i]
        // fmt.Println(len(s)-(i+1))
        // fmt.Println(len(s)-i)
        // fmt.Println(char)
        // 將羅馬字母轉成數字
		num = roman[char]
        // 之前扣
        // fmt.Println("num: ", num)
        // fmt.Println("lastint: ", lastint)
        // 如果 lastint 比較小, total 加上 num
		if num < lastint {
            // 也會出現用相減來表示, 像是 VI 就是 5-1 = 4
            // 只有右邊字母大於左邊字母, 才需要 total - num
			total = total - num
		} else {
            // total 是多個羅馬字母轉換成數字然後相加的結果
			total = total + num
		}

        // fmt.Println("total: ", total)
		lastint = num
	}
	return total
}
// 給定一個羅馬數字，將其轉換成整數。輸入確保在 1 到 3999 的範圍內。
var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}