# [9. Palindrome Number](https://leetcode.com/problems/palindrome-number/)

## 題目

Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

**Example 1**:

```
Input: 121
Output: true
```

**Example 2**:

```
Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

**Example 3**:

```
Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

**Follow up**:

Coud you solve it without converting the integer to a string?

## 題目大意

判斷一個整數是否是回文數。回文數是指正序（從左向右）和倒序（從右向左）讀都是一樣的整數。

## 解題思路

- 判斷一個整數是不是回文數。
- 簡單題。注意會有負數的情況，負數，個位數，10 都不是回文數。其他的整數再按照回文的規則判斷。

## 程式碼

```go

package leetcode

import "strconv"

// 解法一
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	sz := len(arr)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

// 解法二 數字轉字串
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}


```
