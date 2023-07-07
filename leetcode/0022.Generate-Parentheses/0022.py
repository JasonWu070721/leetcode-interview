
class Solution:

    def generateParenthesis(self, n: int) -> List[str]:
        ansList = []
        curStr = ""
        self.dfs(curStr, n, 0, 0, ansList)
        return ansList
        # openCnt 紀錄左括號的數量
        # closeCnt 紀錄右括號的數量s

    def dfs(self, curStr: str, n: int, openCnt: int, closeCnt: int, ansList):
        # 終止條件, 當配對到 n*2 長度的括號, 表示找到一組組合, 將這組合加入 ansList
        if len(curStr) == n*2:
            ansList.append(curStr)
            return
        else:
            if openCnt < n:
                self.dfs(curStr+'(', n, openCnt+1, closeCnt, ansList)
            if closeCnt < openCnt:
                self.dfs(curStr+')', n, openCnt, closeCnt+1, ansList)
