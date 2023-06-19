package leetcode
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for nums_i, nums_v := range nums {
        //fmt.Println("nums_i: ", nums_i)
        //fmt.Println("nums_v: ", nums_v)
        //fmt.Println("target-nums_v: ", target-nums_v)
        // 如果 target-v(表示為差值) 存在 m 裡, 回傳 True, 即找到 target
		if idx, ok := m[target-nums_v]; ok {
            //fmt.Println("idx: ", idx)

            //fmt.Println("ok: ", ok)
			return []int{idx, nums_i}
		}
  
        // 紀錄過的 nums_index 存在  m
        // 紀錄第一個加數在 hash table, 之後可以用 target-nums_v(第二個被加數)
        // 快速查詢 target-nums_v(第二個被加數) 是否存在 m[nums_v(第一個加數)] 內
        // 因為 target-nums_v(第二個被加數) = nums_v(第一個加數)
		m[nums_v] = nums_i
	}
	return nil
}