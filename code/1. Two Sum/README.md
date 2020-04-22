# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 题目

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example

```
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 翻译

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。


## 思路
1 暴力法 直接双重循环遍历每个元素并查找是否存在一个值和target-x 相等的 
2 利用hash 典型的 **空间换时间** 方法

  有人说 比如 java 里面 map.containsKey 方法里面有resize 和 用到了红黑树 准确的来说 一般来说是O(1) 最坏情况是O(n) 没命中才会去遍历红黑树。总的来说hash 还是要比暴力快