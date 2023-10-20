```go
func reverse(x int) int {
    tmp := 0
	for x != 0 {
        //假設 輸入值是 123    第一次 tail = 3 , 第二次 tail = 2 , 第三次 tail = 1
        tail := x%10
        //假設 輸入值是 123   tmp 第一次為 0*10+3 = 3 , 第二次為 3*10+2 = 32 , 第三次為 32*10+1 = 321
		tmp = tmp*10 + tail
		x = x / 10
	}

    //判斷溢出, 如果輸出值大於溢出值
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}
```
