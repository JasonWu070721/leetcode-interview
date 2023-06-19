class Solution:
    def longestPalindrome(self, s: str) -> str:
        # Preventation
        if len(s) <= 1: return s
        
        # Init
        # DP 定義成第i 字母到 第j 字母是否是 Palindromic
        # DP i 和 j, 表示在str 的雙指標
        dp = [[False]*len(s) for _ in range(len(s))]
        # 斜對角都設為 True, dp[i][i], 雙指標指向同一個 index
        for i in range(len(s)): dp[i][i] = True
        
        max_sub_str = s[0]

        # Complete the DP table
        for i in range(len(s)-2, -1, -1):
            for j in range(i+1, len(s)):
                # print("i, j:", i, j)
                if s[i] == s[j]:
                    # i+1 和 j-1 表示 
                    # j-i==1 表示j 在 i 旁邊, 然後 s[i] == s[j], 所以是 Palindromic
                    # dp[i+1][j-1] 表示 i 的右邊一格,和j的右邊一格, 在 DP 表格中屬於 Palindromic
                    # 因為 i 和 j 一開始都在左右旁邊, 距離為0, 所以當距離0的DP為 Palindromic, 當距離為1時也是 Palindromic
                    if j-i==1 or dp[i+1][j-1] == True:
                        dp[i][j] = True
                        
                        # max_sub_str 為最大回文數
                        if len(max_sub_str) < j-i+1:
                            max_sub_str = s[i:j+1]
        
        return max_sub_str