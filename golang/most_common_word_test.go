package leetcode

import "testing"

func TestMostCommonWord(t *testing.T) {
  test := "Bob hit a ball, the hit BALL flew far after it was hit."
  banned := []string{ "hit" }

  word := mostCommonWord(test, banned)

  if word != "ball" {
    t.Errorf("Expect ball but got %v", word)
  }

  test = "a, a, a, a, b,b,b,c, c"
  banned = []string{ "a" }

  word = mostCommonWord(test, banned)


  if word != "b" {
    t.Errorf("Expect b but got %v", word)
  }
}
