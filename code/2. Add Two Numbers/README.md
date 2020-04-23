# [1. Two Sum](https://leetcode.com/problems/add-two-numbers/)

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
 1 自己使用数字加减，把链表编程实际数字，实现了 时间复杂度O(Max(a,b)) 但是 在大数的时候 就会溢出 ，这个没考虑到，使用BigDecimal 解决了 但想想不够巧妙。

 2  如图 对两数相加方法的可视化: 342 + 465 = 807342+465=807，每个结点都包含一个数字，并且数字按位逆序存储。
  ![](https://i.loli.net/2020/04/23/FD6uMT91glGxvsn.jpg)

    从低位到高位遍历两个链表，注意进位的时候注意留到下一个迭代初始化 当前节点为 头结点的前一个节点。

3 通过单链表的定义可以得知，单链表也是递归结构，因此，也可以使用递归的方式来进行reverse操作。

> 由于单链表是线性的，使用递归方式将导致栈的使用也是线性的，当链表长度达到一定程度时，递归会导致爆栈，因此，现实中并不推荐使用递归方式来操作链表。