# [1. Two Sum](https://leetcode.com/problems/two-sum/)


## 思路
1 暴力法 直接双重循环遍历每个元素并查找是否存在一个值和target-x 相等的 
2 利用hash 典型的 **空间换时间** 方法

  有人说 比如 java 里面 map.containsKey 方法里面有resize 和 用到了红黑树 准确的来说 一般来说是O(1) 最坏情况是O(n) 没命中才会去遍历红黑树。总的来说hash 还是要比暴力快

