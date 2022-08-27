package leetcode

import (
  "testing"
)

func TestRob(t *testing.T) {
  nums := []int{ 1, 2, 3, 1 }

  max := rob(nums)

  if max != 4 {
    t.Errorf("Expected 4 but got %v", max)
  }

}
