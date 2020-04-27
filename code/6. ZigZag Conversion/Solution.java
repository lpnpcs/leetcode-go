package zig_zag_conversion6;

import java.util.ArrayList;
import java.util.List;

/**
 * @author lpnpcs
 * @date 2020-04-27 13:55
 */
public class Solution {
    public String convert(String s, int numRows) {
        if(numRows < 2) {
            return s;
        }
        StringBuilder[] rows = new StringBuilder[numRows];
        // 初始化 numRows个 StringBuilder
        for (int i = 0; i < numRows; i++) {
            rows[i] = new StringBuilder();
        }
        // 遍历方向 从上到下是正的 从下到上取负值
        int dir = 1;
        int index = 0;
        for (char c : s.toCharArray()) {
            rows[index].append(c);
            index += dir;
            if (index==0 ||index ==numRows - 1) {
                dir = -dir;
            }
        }
        StringBuilder sb = new StringBuilder();
        for (StringBuilder row : rows) {
            sb.append(row);
        }
        return sb.toString();
    }
}
