package leetcode

import "testing"

func TestStackPushPop(t *testing.T) {
  s := NewStack[int]()

  i := -1
  for i < 1<<10 {
    i++
    s.Push(i)
  }

  for i > -1 {
    v, err := s.Pop()
    if err != nil {
      t.Errorf("Error while popping i=%v: %v", i, err)
      return
    }

    if *v != i {
      t.Errorf("Expect %v but got %v", i, v)
      return
    }

    i--
  }

}
