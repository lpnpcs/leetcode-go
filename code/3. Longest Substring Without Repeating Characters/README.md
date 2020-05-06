# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)


## 思路
 方法一：暴力法 (题目更新后由于时间限制，会出现 TLE)
       逐个检查所有的子字符串，看它是否不含有重复的字符。时间复杂度：O(n^3) 空间复杂度：O(min(n, m))
 方法二: 滑动窗口
  用一个 hashmap 来建立字符和其出现位置之间的映射。

 维护一个滑动窗口，窗口内的都是没有重复的字符，去尽可能的扩大窗口的大小，窗口不停的向右滑动。

（1）如果当前遍历到的字符从未出现过，那么直接扩大右边界；

（2）如果当前遍历到的字符出现过，则缩小窗口（左边索引向右移动），然后继续观察当前遍历到的字符；

（3）重复（1）（2），直到左边索引无法再移动；

（4）维护一个结果 res，每次用出现过的窗口大小来更新结果 res，最后返回 res 获取结果。

