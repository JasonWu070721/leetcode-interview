- 兩兩相鄰的元素，翻轉鏈表

```python
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # 這寫法比較特別, pre 取代 self, pre 就是 self
        # pre.next 表示 self.next
        pre, pre.next = self, head
        # 判斷 pre.next and pre.next.next 兩點是否存在
        while pre.next and pre.next.next:
            # 將下兩個點存到 a, b
            a = pre.next
            b = pre.next.next
            # 將a, b 點交換, 並存到下兩個點
            # pre.next = pre.next.next
            # pre.next.next = pre.next
            # 這是 index 1, 2 交換
            # pre.next = pre.next.next.next, 將pre.next 移到第3個 index

            # 無法理解順序
            pre.next, b.next, a.next = b, a, b.next
            # 此時 pre 已經改成 Optional[ListNode]
            # 將 pre index  移動到 2
            pre = a
        return self.next
```
