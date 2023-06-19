# [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

## 題目

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
```

Explanation: 342 + 465 = 807.

## 題目大意

2 個逆序的鏈表，要求從低位開始相加，得出結果也逆序輸出，返回值是逆序結果鏈表的頭結點。

## 解題思路

需要注意的是各種進位問題。

極端情況，例如

```
Input: (9 -> 9 -> 9 -> 9 -> 9) + (1 -> )
Output: 0 -> 0 -> 0 -> 0 -> 0 -> 1
```

為了處理方法統一，可以先建立一個虛擬頭結點，這個虛擬頭結點的 Next 指向真正的 head，這樣 head 不需要單獨處理，直接 while 循環即可。另外判斷循環終止的條件不用是 p.Next ！= nil，這樣最後一位還需要額外計算，循環終止條件應該是 p != nil。
