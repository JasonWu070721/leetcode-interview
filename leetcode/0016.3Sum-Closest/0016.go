package leetcode

import (
	"math"
	"sort"
)
func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
        // 需要先排序
		sort.Ints(nums)
        // i 為單向指標
		for i := 0; i < n-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
            // j, k 用雙指標
			for j, k := i+1, n-1; j < k; {
                // 計算 i, j, k 的合
				sum := nums[i] + nums[j] + nums[k]
                // diff 預設是 MaxInt32
                // abs(sum-target) 是計算 sum 和 target 的差值
                // abs(sum-target) < diff 是計算最小的 diff
				if abs(sum-target) < diff {
                    // diff 會計算出最小的 abs(sum-target)
                    // 這裡會取得最小的 diff, 順便紀錄 sum 做為 res
					res, diff = sum, abs(sum-target)
				}

                // 如果直接找到 target 就 return res
                // 如果 sum 比較大, k 向左移動
                // 如果 sum 比較小, j 向右移動
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}