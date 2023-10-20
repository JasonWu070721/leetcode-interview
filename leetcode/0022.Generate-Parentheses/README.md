- 給出 n 代表生成括號的對數，請你寫出一個函數，使其能夠生成所有可能的並且有效的括號組合。

```python

class Solution:

    def generateParenthesis(self, n: int) -> List[str]:
        ansList = []
        curStr=""
        self.dfs(curStr,n,0,0,ansList)
        return ansList
		# openCnt 紀錄左括號的數量
		# closeCnt 紀錄右括號的數量
    def dfs(self,curStr:str,n:int,openCnt:int,closeCnt:int,ansList:[]):
        # 終止條件, 當配對到 n*2 長度的括號, 表示找到一組組合, 將這組合加入 ansList
        if len(curStr) == n*2:
            ansList.append(curStr)
            return
        else:
            if openCnt < n:
                self.dfs(curStr+'(',n,openCnt+1,closeCnt,ansList)
            if closeCnt<openCnt:
                self.dfs(curStr+')',n,openCnt,closeCnt+1,ansList)

```
