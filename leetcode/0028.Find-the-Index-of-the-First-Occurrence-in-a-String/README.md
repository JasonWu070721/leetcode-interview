- 實現一個尋找 substring 的函數。如果在母串中找到了子串，返回子串在母串中出現的下標，如果沒有找到，返回 -1，如果子串是空串，則返回 0

```go
func strStr(haystack string, needle string) int {
    for i := 0; ; i++ {
        for j := 0; ; j++ {
            // 找到相同的字串, 回傳遞一個字的 index
            // 此時 i 停在 haystack 中的 needle 的開始位置
            // 因為 needle[j] == haystack[i+j] 就會繼續,
            // 所以當 j == len(needle), 表示 needle 全都在 haystack 找到
            // 表示所有 needle 都比對完成, 回傳字母開頭 i
            if j == len(needle) {
                return i
            }

            // i 會一直向右邊移動, 直到 i + 部分的 needle, 等於 len(haystack)
            // 因為 最後面只有部分的 needle 相同, 所以在移動 i, 就可以提前結束
            // i 是比對過程中 needle 的前一個字母, 如果 j 已經是 needle 最後一個字母,
            // 表示已經比對到 haystack 最後一個字母
    		if i+j == len(haystack) {
    			return -1
    		}

    		// j 表示 needle 的 index,
    		// needle[j] != haystack[i+j] 表示沒找到, 跳出 for j := 0; ; j++,
            // i + j 表示在 haystack 移動到 i 的位置, 再加上要找的 needle 中的 j 位置
            // needle[j] == haystack[i+j] 表示找到一個字
            // i 表示開始做比對的起點
            // j 表示 needle 要比對的 index
    		if needle[j] != haystack[i+j] {
    			break
    		}
    	}
    }
}
```
