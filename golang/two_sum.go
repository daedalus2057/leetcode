package leetcode
// @leetup=custom
// @leetup=info id=1 lang=golang slug=two-sum

/*
* Given an array of integers `nums` and an integer `target`, return *indices of
* the two numbers such that they add up to `target`*.
* 
* You may assume that each input would have *exactly* one solution, and you
* may not use the *same* element twice.
* 
* You can return the answer in any order.
* 
* 
* Example 1:
* 
* Input: nums = [2,7,11,15], target = 9
* Output: [0,1]
* Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
* 
* Example 2:
* 
* Input: nums = [3,2,4], target = 6
* Output: [1,2]
* 
* Example 3:
* 
* Input: nums = [3,3], target = 6
* Output: [0,1]
* 
* 
* Constraints:
* 
* * `2 <= nums.length <= 104`
* * `-109 <= nums[i] <= 109`
* * `-109 <= target <= 109`
* * Only one valid answer exists.
* 
* 
* Follow-up: Can you come up with an algorithm that is less than `O(n2) `time
* complexity?
* 
*/
// @leetup=custom
// @leetup=code
import "sort"

func twoSum(nums []int, target int) []int {
  for i, a := range nums[:len(nums) - 1] {
    for j, b := range nums[i+1:] {
      if a + b == target {
        return []int{ i, (j+i+1) }
      }
    }
  }

  return []int{ -1, -1, }
}

func twoSumNLogN(nums []int, target int) []int {
  sort.Ints(nums)

  i := 0
  j := len(nums) - 1
  mid := len(nums) / 2
  cand := nums[mid]
  for i < j {
    if target == cand {
      break
    }

    if target < cand {
      j = mid - 1
      mid = (j - i) / 2
      cand = nums[mid]
      continue
    }

    i = mid + 1
    mid = (j - 1) / 2
    cand = nums[mid]
  }
  
  //mid = j / 2
  //i = 0
  // idea: start from mid and move out from mid - 1 to the left
  // and mid to the right until left is 0 or right is j
  // compare sums to target.

  // proposed runtime n log n dominates the algo
  return []int{ 0 }

}
// @leetup=code
