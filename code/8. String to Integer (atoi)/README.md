# [8. String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/)


## 思路
1 按照题目 一步一步写，但是这个题test case 写的比较变态，很难一次作对
2 使用自动机 这里引用官方图片

![自动机](https://assets.leetcode-cn.com/solution-static/8_fig1.PNG)

|           | **' '**   | **+/-** | **number** | **other** |
| --------- | ----- | ------- | ---------- | --------- |
| start     | start | signed  | in_number  | end       |
| signed    | end   | end     | in_number  | end       |
| in_number | end   | end     | in_number  | end       |
| end       | end   | end     | end        | end       |

> https://leetcode-cn.com/problems/string-to-integer-atoi/solution/javascriptzi-dong-ji-guan-fang-ti-jie-de-xiang-xi-/