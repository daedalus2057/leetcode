package leetcode
// @leetup=custom
// @leetup=info id=215 lang=golang slug=kth-largest-element-in-an-array

/*
* Given an integer array `nums` and an integer `k`, return *the* `kth` *largest
* element in the array*.
* 
* Note that it is the `kth` largest element in the sorted order, not the `kth`
* distinct element.
* 
* You must solve it in `O(n)` time complexity.
* 
* 
* Example 1:
* 
* Input: nums = [3,2,1,5,6,4], k = 2
* Output: 5
* 
* Example 2:
* 
* Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
* Output: 4
* 
* 
* Constraints:
* 
* * `1 <= k <= nums.length <= 105`
* * `-104 <= nums[i] <= 104`
* 
*/
// @leetup=custom
// @leetup=code

func findKthLargest(nums []int, k int) int {

  // find the min and max values
  min := 10_001
  max := -10_001
  for _, v := range nums {
    if v < min {
      min = v
    }

    if v > max {
      max = v
    }
  }

  n := (max - min) + 1

  linearSearchArray := make([]int, (2*n) + 1)
  for _, x := range nums {
    linearSearchArray[x+n] += 1
  }

  j := len(linearSearchArray) - 1
  kth := 0
  for j > -1 && kth < k {
    if linearSearchArray[j] > 0 {
      kth += linearSearchArray[j]
    }

    j--
  }

  return j+1 - n
}
// @leetup=code
