package leetcode
func maxArea(height []int) int {
    // 雙指標方式, 計算最大面積, 從最左邊和最右邊往中間計算
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
        // 判斷左邊高還是右邊高
		if height[start] < height[end] {
            // 如果左邊高用左邊高度
			high = height[start]
            // 計算完往右邊移動一格
			start++
		} else {
            // 如果右邊高用右邊高度
			high = height[end]
            // 計算完往左邊移動一格
			end--
		}

        // 計算最大面積
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}