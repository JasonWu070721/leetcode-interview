# [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

## 題目

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg)

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 1:

```c
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
```

## 題目大意

給出一個非負整數陣列 a1，a2，a3，…… an，每個整數標識一個豎立在坐標軸 x 位置的一堵高度為 ai 的牆，選擇兩堵牆，和 x 軸構成的容器可以容納最多的水。

## 解題思路

這一題也是對撞指針的思路。首尾分別 2 個指針，每次移動以後都分別判斷長寬的乘積是否最大。
