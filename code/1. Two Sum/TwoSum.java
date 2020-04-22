import java.util.HashMap;
import java.util.Map;

/**
 * @author lpnpcs
 * @date 2020-04-22 15:39
 */
public class TwoSum {
    public  int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>(16);
        for (int i = 0; i < nums.length; i++) {
            int c = target - nums[i];
            if (map.containsKey(c)) {
                return new int[]{
                        map.get(c), i
                };
            }
            map.put(nums[i], i);
        }
        throw new IllegalArgumentException("No two sum solution");
    }

    public static void main(String[] args) {
        new TwoSum().twoSum(new int[]{2, 7, 11, 5}, 9);
    }
}
