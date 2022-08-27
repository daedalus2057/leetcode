package leetcode

import "testing"

func TestLongestPalSub(t *testing.T) {
  lp := longestPalindrome("ccc")
  if lp != "ccc" {
    t.Errorf("Expect 'ccc' but got '%v'", lp)
  }

  lp = longestPalindrome("babb")
  if lp != "bab" {
    t.Errorf("Expect 'bab' but got '%v'", lp)
  }
}
