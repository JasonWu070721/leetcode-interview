# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:

        dummy = head = ListNode(0)
        # 需要紀錄進位是多少
        carry = 0
        # 當l1 and l2 都存在時
        while(l1 and l2):
            # 計算兩個值相加跟進位數
            this_val = l1.val + l2.val + carry
            # 判斷是否會進位
            carry = this_val // 10
            # 寫入相加結果, %10 是移除進位數, 下一個迴圈會加上 carry 
            head.next = ListNode(this_val % 10)
            print(head.next)
            l1 = l1.next
            l2 = l2.next
            head = head.next
        else:
            # 當l1 存在, l2不存在
            while(l1):
                # 只處理l1 和進位數, 可能之前 l1.val + l2.val 存在進位數
                this_val = l1.val + carry
                carry = this_val // 10            
                head.next = ListNode(this_val % 10)
                l1 = l1.next
                head = head.next
            # 當l2 存在, l1不存在
            while(l2):
                this_val = l2.val + carry
                carry = this_val // 10            
                head.next = ListNode(this_val % 10)
                l2 = l2.next
                head = head.next
            if carry:
                head.next = ListNode(carry)

        return dummy.next