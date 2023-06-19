# [8. String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/)

## 題目

Implement the `myAtoi(string s)` function, which converts a string to a 32-bit signed integer (similar to C/C++'s `atoi` function).

The algorithm for `myAtoi(string s)` is as follows:

1. Read in and ignore any leading whitespace.
2. Check if the next character (if not already at the end of the string) is `'-'` or `'+'`. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
3. Read in next the characters until the next non-digit charcter or the end of the input is reached. The rest of the string is ignored.
4. Convert these digits into an integer (i.e. `"123" -> 123`, `"0032" -> 32`). If no digits were read, then the integer is `0`. Change the sign as necessary (from step 2).
5. If the integer is out of the 32-bit signed integer range `[-231, 231 - 1]`, then clamp the integer so that it remains in the range. Specifically, integers less than `231` should be clamped to `231`, and integers greater than `231 - 1` should be clamped to `231 - 1`.
6. Return the integer as the final result.

**Note:**

- Only the space character `' '` is considered a whitespace character.
- **Do not ignore** any characters other than the leading whitespace or the rest of the string after the digits.

**Example 1:**

```
Input: s = "42"
Output: 42
Explanation: The underlined characters are what is read in, the caret is the current reader position.
Step 1: "42" (no characters read because there is no leading whitespace)
         ^
Step 2: "42" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "42" ("42" is read in)
           ^
The parsed integer is 42.
Since 42 is in the range [-231, 231 - 1], the final result is 42.

```

**Example 2:**

```
Input: s = "   -42"
Output: -42
Explanation:
Step 1: "   -42" (leading whitespace is read and ignored)
            ^
Step 2: "   -42" ('-' is read, so the result should be negative)
             ^
Step 3: "   -42" ("42" is read in)
               ^
The parsed integer is -42.
Since -42 is in the range [-231, 231 - 1], the final result is -42.

```

**Example 3:**

```
Input: s = "4193 with words"
Output: 4193
Explanation:
Step 1: "4193 with words" (no characters read because there is no leading whitespace)
         ^
Step 2: "4193 with words" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "4193 with words" ("4193" is read in; reading stops because the next character is a non-digit)
             ^
The parsed integer is 4193.
Since 4193 is in the range [-231, 231 - 1], the final result is 4193.

```

**Example 4:**

```
Input: s = "words and 987"
Output: 0
Explanation:
Step 1: "words and 987" (no characters read because there is no leading whitespace)
         ^
Step 2: "words and 987" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "words and 987" (reading stops immediately because there is a non-digit 'w')
         ^
The parsed integer is 0 because no digits were read.
Since 0 is in the range [-231, 231 - 1], the final result is 0.

```

**Example 5:**

```
Input: s = "-91283472332"
Output: -2147483648
Explanation:
Step 1: "-91283472332" (no characters read because there is no leading whitespace)
         ^
Step 2: "-91283472332" ('-' is read, so the result should be negative)
          ^
Step 3: "-91283472332" ("91283472332" is read in)
                     ^
The parsed integer is -91283472332.
Since -91283472332 is less than the lower bound of the range [-231, 231 - 1], the final result is clamped to -231 = -2147483648.
```

**Constraints:**

- `0 <= s.length <= 200`
- `s` consists of English letters (lower-case and upper-case), digits (`0-9`), `' '`, `'+'`

## 題目大意

請你來實現一個 myAtoi(string s) 函數，使其能將字串轉換成一個 32 位有符號整數（類似 C/C++ 中的 atoi 函數）。

函數 myAtoi(string s) 的演算法如下：

- 讀入字串並丟棄無用的前導空格
- 檢查下一個字元（假設還未到字元末尾）為正還是負號，讀取該字元（如果有）。 確定最終結果是負數還是正數。 如果兩者都不存在，則假定結果為正。
- 讀入下一個字元，直到到達下一個非數字字元或到達輸入的結尾。字串的其餘部分將被忽略。
- 將前面步驟讀入的這些數字轉換為整數（即，"123" -> 123， "0032" -> 32）。如果沒有讀入數字，則整數為 0 。必要時更改符號（從步驟 2 開始）。
- 如果整數數超過 32 位有符號整數範圍 [−231, 231 − 1] ，需要截斷這個整數，使其保持在這個範圍內。具體來說，小於 −231 的整數應該被固定為 −231 ，大於 231 − 1 的整數應該被固定為 231 − 1 。
- 返回整數作為最終結果。

注意：

- 本題中的空白字元只包括空格字元 ' ' 。
- 除前導空格或數字後的其餘字串外，請勿忽略 任何其他字元。

## 解題思路

- 這題是簡單題。題目要求實現類似 `C++` 中 `atoi` 函數的功能。這個函數功能是將字串類型的數字轉成 `int` 類型數字。先去除字串中的前導空格，並判斷記錄數字的符號。數字需要去掉前導 `0` 。最後將數字轉換成數字類型，判斷是否超過 `int` 類型的上限 `[-2^31, 2^31 - 1]`，如果超過上限，需要輸出邊界，即 `-2^31`，或者 `2^31 - 1`。

## 程式碼

```go
package leetcode

func myAtoi(s string) int {
	maxInt, signAllowed, whitespaceAllowed, sign, digits := int64(2<<30), true, true, 1, []int{}
	for _, c := range s {
		if c == ' ' && whitespaceAllowed {
			continue
		}
		if signAllowed {
			if c == '+' {
				signAllowed = false
				whitespaceAllowed = false
				continue
			} else if c == '-' {
				sign = -1
				signAllowed = false
				whitespaceAllowed = false
				continue
			}
		}
		if c < '0' || c > '9' {
			break
		}
		whitespaceAllowed, signAllowed = false, false
		digits = append(digits, int(c-48))
	}
	var num, place int64
	place, num = 1, 0
	lastLeading0Index := -1
	for i, d := range digits {
		if d == 0 {
			lastLeading0Index = i
		} else {
			break
		}
	}
	if lastLeading0Index > -1 {
		digits = digits[lastLeading0Index+1:]
	}
	var rtnMax int64
	if sign > 0 {
		rtnMax = maxInt - 1
	} else {
		rtnMax = maxInt
	}
	digitsCount := len(digits)
	for i := digitsCount - 1; i >= 0; i-- {
		num += int64(digits[i]) * place
		place *= 10
		if digitsCount-i > 10 || num > rtnMax {
			return int(int64(sign) * rtnMax)
		}
	}
	num *= int64(sign)
	return int(num)
}
```
