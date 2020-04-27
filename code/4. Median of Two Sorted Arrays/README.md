# 4. [Median of Two Sorted Arrayss](https://leetcode.com/problems/median-of-two-sorted-arrays/)

## 题目

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5



## 思路
1 简单粗暴，先将两个数组合并，两个有序数组的合并也是归并排序中的一部分。然后根据奇数，还是偶数，返回中位数。
  时间复杂度：遍历全部数组 (m+n)(m+n)
  空间复杂度：开辟了一个数组，保存合并后的两个数组 O(m+n)O(m+n)  
  不满足题意
2 比较普适的办法是 因为中位数 是在 两个数组的中间 可以抽象为求两个有序数组的第k 小的数字 这个算法比较有普适性，每次去掉k/2 个数字 使用尾递归 不会增额外的栈空间 
  注意 为什么 k / 2 小的那个 一定不包含第k 小？ 
    因为大的数 已经是 目前第k 小的数字了 小的 里面所有的都比这些小 所以肯定不是。
   
  > https://leetcode-cn.com/problems/search-insert-position/solution/te-bie-hao-yong-de-er-fen-cha-fa-fa-mo-ban-python-/
  > https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/he-bing-yi-hou-zhao-gui-bing-guo-cheng-zhong-zhao-/
