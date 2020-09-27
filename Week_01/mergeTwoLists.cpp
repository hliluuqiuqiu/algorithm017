/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
*/


/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode *ret = new ListNode();
        ListNode *cusor = ret;
        while (l1 != NULL && l2 != NULL)
        {
            if (l1->val <= l2->val) {
                cusor->next = l1;
                cusor = l1;
                l1 = l1->next;
            } else {
                cusor->next = l2;
                cusor = l2;
                l2 = l2->next;
            }
        }

        cusor->next = (l1 == NULL ? l2 : l1);
        
        return ret->next;
    }
};