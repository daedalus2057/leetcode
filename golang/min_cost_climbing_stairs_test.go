package leetcode

import "testing"


func TestMinCostStairs(t *testing.T) {
  nums := []int{ 10, 15, 20 }

  cost := minCostClimbingStairs(nums)

  if cost != 15 {
    t.Errorf("Expect 15 but got %v", cost)
  }

}
