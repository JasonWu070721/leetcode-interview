package leetcode
func isPalindrome(x int) bool {
	//  如果是負數就 return false
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	// 因為最右邊不會是0, 有0 無法形成 Palindrome 
	if x%10 == 0 {
		return false
	}

	// 將 int 轉成  arr = [1, 2, 1]
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}

	sz := len(arr)
	// fmt.Println("sz", sz)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}