package leetcode

import "testing"

func TestKthLargest(t *testing.T) {
  nums := []int{ 3,2,1,5,6,4 }

  kth := findKthLargest(nums, 2)

  if kth != 5 {
    t.Errorf("Expect 5 but got %v", kth)
  }

  nums = []int{ 3,2,3,1,2,4,5,5,6 }

  kth = findKthLargest(nums, 4)

  if kth != 4 {
    t.Errorf("Expect 4 but got %v", kth)
  }
}
