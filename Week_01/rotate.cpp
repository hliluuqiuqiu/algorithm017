
/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。
*/
class Solution {
public:
    void rotate(vector<int>& nums, int k) {
        int round = k % nums.size();
        if (nums.size() <= 1 || round == 0)
            return;
        int count = 0;
        for (int i = 0; count < nums.size(); i++) {
            int curent = i;
            int prev = nums[curent];
            do {
                int next = (curent + round) % nums.size();
                int temp = nums[next];
                nums[next] = prev;
                prev = temp;
                curent = next;
                count++;
            } while (curent != i);
        }
    }
};