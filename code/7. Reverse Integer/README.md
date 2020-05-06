# [7. Reverse Integer](https://leetcode.com/problems/reverse-integer/)


## 思路

1 直觉 先转成字符串，然后 翻转 转数字。 用到库 不太好。
2 用数字 方法 弹入 弹出 ，并且每次判断边界值
  ```
  //pop operation:
pop = x % 10;
x /= 10;

//push operation:
temp = rev * 10 + pop;
rev = temp;
  ```
