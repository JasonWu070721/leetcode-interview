class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        # 先對nums排序
        nums = sorted(nums)
        # print(nums)

        result = set()
        for i in range(len(nums)):
            # 用雙指標, l 是左指標, r 是右指標
            l = i + 1
            r = len(nums) - 1
            # 因為  nums[l] + nums[r] + target = 0,  nums[l] + nums[r] = - target
            # 所以先讓 target 變成負數
            target = 0 - nums[i]
            while l < r:
                # 如果 l, r 的值 == target, 表示找到  nums[l] + nums[r] +  target = 0
                if nums[l] + nums[r] == target:
                    # 新增找到的值到 result
                    result.add((nums[i], nums[l], nums[r]))
                    # l 往右走一步, r 往左走一步
                    l += 1
                    r -= 1
                # 因為已經排序過, 所以發現 nums[l] + nums[r] < target,
                # l 往右走一步, 讓 nums[l] + nums[r] 變大, 更接近target
                elif nums[l] + nums[r] < target:
                    # l 往右走一步
                    l += 1
                else:
                    # r 往左走一步
                    r -= 1
        return list(result)
