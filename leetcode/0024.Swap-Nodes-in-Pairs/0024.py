class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        pre, pre.next = self, head
        # 判斷是否來有兩個點
        while pre.next and pre.next.next:
            # 將下兩個點存到 a, b
            a = pre.next
            b = a.next
            # 將a, b 點交換, 並存到下兩個點
            pre.next, b.next, a.next = b, a, b.next
            pre = a
        return self.next
