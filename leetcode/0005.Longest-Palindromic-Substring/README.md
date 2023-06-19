# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## 題目

Given a string `s`, return _the longest palindromic substring_ in `s`.

**Example 1:**

```
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.

```

**Example 2:**

```
Input: s = "cbbd"
Output: "bb"

```

**Example 3:**

```
Input: s = "a"
Output: "a"

```

**Example 4:**

```
Input: s = "ac"
Output: "a"

```

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters (lower-case and/or upper-case),

## 題目大意

給你一個字串 `s`，找到 `s` 中最長的回文子串。

## 解題思路

- 此題非常經典，並且有多種解法。
- 解法一，動態規劃。定義 `dp[i][j]` 表示從字串第 `i` 個字元到第 `j` 個字元這一段子串是否是回文串。由回文串的性質可以得知，回文串去掉一頭一尾相同的字元以後，剩下的還是回文串。所以狀態轉移方程是 `dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])`，注意特殊的情況，`j - i == 1` 的時候，即只有 2 個字元的情況，只需要判斷這 2 個字元是否相同即可。`j - i == 2` 的時候，即只有 3 個字元的情況，只需要判斷除去中心以外對稱的 2 個字元是否相等。每次循環動態維護保存最長回文串即可。時間複雜度 O(n^2)，空間複雜度 O(n^2)。
- 解法二，中心擴散法。動態規劃的方法中，我們將任意起始，終止範圍內的字串都判斷了一遍。其實沒有這個必要，如果不是最長回文串，無需判斷並保存結果。所以動態規劃的方法在空間複雜度上還有最佳化空間。判斷回文有一個核心問題是找到“軸心”。如果長度是偶數，那麼軸心是中心虛擬的，如果長度是奇數，那麼軸心正好是正中心的那個字母。中心擴散法的思想是列舉每個軸心的位置。然後做兩次假設，假設最長回文串是偶數，那麼以虛擬中心往 2 邊擴散；假設最長回文串是奇數，那麼以正中心的字元往 2 邊擴散。擴散的過程就是對稱判斷兩邊字元是否相等的過程。這個方法時間複雜度和動態規劃是一樣的，但是空間複雜度降低了。時間複雜度 O(n^2)，空間複雜度 O(1)。
- 解法三，滑動窗口。這個寫法其實就是中心擴散法變了一個寫法。中心擴散是依次列舉每一個軸心。滑動窗口的方法稍微最佳化了一點，有些軸心兩邊字元不相等，下次就不會列舉這些不可能形成回文子串的軸心了。不過這點最佳化並沒有最佳化時間複雜度，時間複雜度 O(n^2)，空間複雜度 O(1)。
- 解法四，馬拉車演算法。這個演算法是本題的最優解，也是最複雜的解法。時間複雜度 O(n)，空間複雜度 O(n)。中心擴散法有 2 處有重複判斷，第一處是每次都往兩邊擴散，不同中心擴散多次，實際上有很多重複判斷的字元，能否不重複判斷？第二處，中心能否跳躍選擇，不是每次都列舉，是否可以利用前一次的資訊，跳躍選擇下一次的中心？馬拉車演算法針對重複判斷的問題做了最佳化，增加了一個輔助陣列，將時間複雜度從 O(n^2) 最佳化到了 O(n)，空間換了時間，空間複雜度增加到 O(n)。

  ![https://img.halfrost.com/Leetcode/leetcode_5_1.png](https://img.halfrost.com/Leetcode/leetcode_5_1.png)

- 首先是預處理，向字串的頭尾以及每兩個字元中間新增一個特殊字元 `#`，比如字串 `aaba` 處理後會變成 `#a#a#b#a#`。那麼原先長度為偶數的回文字串 `aa` 會變成長度為奇數的回文字串 `#a#a#`，而長度為奇數的回文字串 `aba` 會變成長度仍然為奇數的回文字串 `#a#b#a#`，經過預處理以後，都會變成長度為奇數的字串。**注意這裡的特殊字元不需要是沒有出現過的字母，也可以使用任何一個字元來作為這個特殊字元。**這是因為，當我們只考慮長度為奇數的回文字串時，每次我們比較的兩個字元奇偶性一定是相同的，所以原來字串中的字元不會與插入的特殊字元互相比較，不會因此產生問題。**預處理以後，以某個中心擴散的步數和實際字串長度是相等的。**因為半徑裡面包含了插入的特殊字元，又由於左右對稱的性質，所以擴散半徑就等於原來回文子串的長度。

  ![https://img.halfrost.com/Leetcode/leetcode_5_2.png](https://img.halfrost.com/Leetcode/leetcode_5_2.png)

- 核心部分是如何通過左邊已經掃描過的資料推出右邊下一次要擴散的中心。這裡定義下一次要擴散的中心下標是 `i`。如果 `i` 比 `maxRight` 要小，只能繼續中心擴散。如果 `i` 比 `maxRight` 大，這是又分為 3 種情況。三種情況見上圖。將上述 3 種情況總結起來，就是 ：`dp[i] = min(maxRight-i, dp[2*center-i])`，其中，`mirror` 相對於 `center` 是和 `i` 中心對稱的，所以它的下標可以計算出來是 `2*center-i`。更新完 `dp[i]` 以後，就要進行中心擴散了。中心擴散以後動態維護最長回文串並相應的更新 `center`，`maxRight`，並且記錄下原始字串的起始位置 `begin` 和 `maxLen`。

## 程式碼

```go
package leetcode

// 解法一 Manacher's algorithm，時間複雜度 O(n)，空間複雜度 O(n)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	newS := make([]rune, 0)
	newS = append(newS, '#')
	for _, c := range s {
		newS = append(newS, c)
		newS = append(newS, '#')
	}
	// dp[i]:    以預處理字串下標 i 為中心的回文半徑(奇數長度時不包括中心)
	// maxRight: 通過中心擴散的方式能夠擴散的最右邊的下標
	// center:   與 maxRight 對應的中心字元的下標
	// maxLen:   記錄最長回文串的半徑
	// begin:    記錄最長回文串在起始串 s 中的起始下標
	dp, maxRight, center, maxLen, begin := make([]int, len(newS)), 0, 0, 1, 0
	for i := 0; i < len(newS); i++ {
		if i < maxRight {
			// 這一行程式碼是 Manacher 演算法的關鍵所在
			dp[i] = min(maxRight-i, dp[2*center-i])
		}
		// 中心擴散法更新 dp[i]
		left, right := i-(1+dp[i]), i+(1+dp[i])
		for left >= 0 && right < len(newS) && newS[left] == newS[right] {
			dp[i]++
			left--
			right++
		}
		// 更新 maxRight，它是遍歷過的 i 的 i + dp[i] 的最大者
		if i+dp[i] > maxRight {
			maxRight = i + dp[i]
			center = i
		}
		// 記錄最長回文子串的長度和相應它在原始字串中的起點
		if dp[i] > maxLen {
			maxLen = dp[i]
			begin = (i - maxLen) / 2 // 這裡要除以 2 因為有我們插入的輔助字元 #
		}
	}
	return s[begin : begin+maxLen]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 解法二 滑動窗口，時間複雜度 O(n^2)，空間複雜度 O(1)
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// 移動到相同字母的最右邊（如果有相同字母）
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 找到回文的邊界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重設到下一次尋找回文的中心
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}

// 解法三 中心擴散法，時間複雜度 O(n^2)，空間複雜度 O(1)
func longestPalindrome2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}

// 解法四 DP，時間複雜度 O(n^2)，空間複雜度 O(n^2)
func longestPalindrome3(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
```
