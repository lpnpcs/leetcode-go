# [9. Palindrome Number](https://leetcode.com/problems/palindrome-number/)


## 思路
1 第一个想法是数字转字符串，判断字符串是否回文，但是问题描述中不允许这么做
2 反转一半数字 先处理临界情况 负数不是回文。然后翻转后半部分(参考第七题使用数学取余取模)