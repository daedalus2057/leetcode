package leetcode
// @leetup=custom
// @leetup=info id=740 lang=golang slug=delete-and-earn

/*
* You are given an integer array `nums`. You want to maximize the number of points
* you get by performing the following operation any number of times:
* 
* * Pick any `nums[i]` and delete it to earn `nums[i]` points. Afterwards, you
*   must delete every element equal to `nums[i] - 1` and every element equal
*   to `nums[i] + 1`.
* 
* Return *the maximum number of points you can earn by applying the above
* operation some number of times*.
* 
* 
* Example 1:
* 
* Input: nums = [3,4,2]
* Output: 6
* Explanation: You can perform the following operations:
* - Delete 4 to earn 4 points. Consequently, 3 is also deleted. nums = [2].
* - Delete 2 to earn 2 points. nums = [].
* You earn a total of 6 points.
* 
* Example 2:
* 
* Input: nums = [2,2,3,3,3,4]
* Output: 9
* Explanation: You can perform the following operations:
* - Delete a 3 to earn 3 points. All 2's and 4's are also deleted. nums = [3,3].
* - Delete a 3 again to earn 3 points. nums = [3].
* - Delete a 3 once more to earn 3 points. nums = [].
* You earn a total of 9 points.
* 
* 
* Constraints:
* 
* * `1 <= nums.length <= 2 * 104`
* * `1 <= nums[i] <= 104`
* 
*/
// @leetup=custom
// @leetup=code

//import "sort"


func deleteAndEarn(nums []int) int {
  // we can memioize this I assume -- for each entry we have a value such that
  // we can have all the numbers that equal this number and NOT have those that
  // are +/-1 we could generate a count of each number and store the results and
  // then we could consider the cost at each step. We could sort the numbers first
  // and start from the top which would make finding the +/- easier

  // now store a count of each value
  // foundValues := make([]int, 0)
  // totalValues := make(map[int]int)
  // for _, v := range nums {
  //   if tot, ok := totalValues[v]; ok {
  //     totalValues[v] = tot + v
  //     continue
  //   }
  //
  //   foundValues = append(foundValues, v)
  //   totalValues[v] = v
  // }

  // the max cost at each step is the max cost so far
  // one more and one less you have A B C if you choose A you
  // remove B if you choose B you remove A and C and if you choose
  // assessLostValue := func(v int, m map[int]int) int {
  //   lost := 0
  //   if tot, ok := m[v+1]; ok {
  //     lost = lost + tot
  //   }
  //
  //   if tot, ok := m[v-1]; ok {
  //     lost = lost + tot
  //   }
  //
  //   return lost
  // }

  // // best value for list of size 1
  // best := 0
  // for value := range foundValues {
  //   if totalValues[value] > assessLostValue(value, totalValues) {
  //         
  //   }
  //
  // }

  // TODO
  return 0
    
}
// @leetup=code
