class Solution:
    def removeDuplicates(self, nums):
        if not nums: return 0
        i, j = 0, 1
        # 終止條件, j < nums 長度
        while j < len(nums):
            # 判斷i 跟i的下一個數字是否相同
            # 如果i 不等於 j
            if nums[i] != nums[j]:
                # i 右走一步
                i += 1
                # i 被 j 取代
                # i 
                nums[i] = nums[j]
            # j 走一步
            j += 1
        # return i 走一步
        return i + 1