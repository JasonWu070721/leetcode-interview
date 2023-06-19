package leetcode
func lengthOfLongestSubstring(s string) int {
    // 解法三 滑動窗口-雜湊桶
   	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for left < len(s) {
        // idx, ok := indexes[s[left]] 表示 s[left]的字母出現在 indexes 內, ok 就為 True
        // right 就要移動到 indexes[s[left]] 的 idx 右邊
        // 避免 right 往回走, 規範 idx >= right, 讓 idx 只能 >= right
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
        // left 是一直向右++
        // s[left] 是 val, left 是 index
        // 紀錄 s[left] 的位置
		indexes[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}