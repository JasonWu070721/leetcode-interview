- 編寫一個函數來尋找字串陣列中的最長公共前綴。

```go
func longestCommonPrefix(strs []string) string {
    // prefix 是比較對象
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
        // i 指單字
		for j := 0; j < len(prefix); j++ {
            // j 表示字母
            // fmt.Println("string(prefix[j]): ", string(prefix[j]))
            // fmt.Println("i, j:", i, j)
            // fmt.Println("strs[i]:", strs[i])
            // len(strs[i])是第i文字的長度, j 是要比較的字串長度, j 不能超過 len(strs[i])
            // strs[i][j] != prefix[j], 表示第j個字不一樣
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
                // fmt.Println("prefix[0:j]: ", prefix[0:j])
                // prefix 為相同的部分, 如果第一個字母就不同 j = 0, prefix = 空字串
				prefix = prefix[0:j]
                // 比到不一樣的字母, 就結束圈, 比下一個文字
				break
			}
		}
	}

	return prefix
}
```
