package leetcode

import (
  "testing"
)


func TestSimple(t *testing.T) {
  grid := [][]byte{
    {'1','1','1','1','0'},
    {'1','1','0','1','0'},
    {'1','1','0','0','0'},
    {'0','0','0','0','0'},
  }

  count, root_a := numIslands(grid)

  if count != 1 {
    t.Errorf("Expected count to be 1 but got %v\nroot_a = %v\n ", count, root_a)
  }

  grid = [][]byte{
    {'1','1','0','0','0'},
    {'1','1','0','0','0'},
    {'0','0','1','0','0'},
    {'0','0','0','1','1'},
  }

  count, root_a = numIslands(grid)

  if count != 3 {
    t.Errorf("Expected count to be 3 but got %v\nroot_a = %v\n ", count, root_a)
  }


  grid = [][]byte{{'1','1','1'},{'0','1','0'},{'1','1','1'},}

  count, root_a = numIslands(grid)

  if count != 1 {
    t.Errorf("Expected count to be 1 but got %v\nroot_a = %v\n ", count, root_a)
  }

}
