- 給定一個有序陣列 nums，對陣列中的元素進行去重，使得原陣列中的每個元素只有一個。最後返回去重以後陣列的長度值。
- 題目要求將一個已排序陣列中所有的重複數字刪去，使得這個陣列的每個數字只出現一次，並且必須要以 in-place 的方式來處理。最終處理過後，回傳新陣列的長度。

```python
class Solution:
    def removeDuplicates(self, nums):
        if not nums: return 0
        i, j = 0, 1
        # i 是連續數字 count, 不重複
        # j 走一回 nums 的 index
        # 終止條件, j < nums 長度
        while j < len(nums):
            # 判斷i 跟i的下一個數字是否相同
            # 如果i 不等於 j
            if nums[i] != nums[j]:
                # i 右走一步
                i += 1
                # i 被 j 取代
                nums[i] = nums[j]
            # j 走一步
            j += 1
        # return i 走一步
        return i + 1
```
