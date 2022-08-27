package leetcode
// @leetup=custom
// @leetup=info id=5 lang=golang slug=longest-palindromic-substring

/*
* Given a string `s`, return *the longest palindromic substring* in `s`.
* 
* 
* Example 1:
* 
* Input: s = "babad"
* Output: "bab"
* Explanation: "aba" is also a valid answer.
* 
* Example 2:
* 
* Input: s = "cbbd"
* Output: "bb"
* 
* 
* Constraints:
* 
* * `1 <= s.length <= 1000`
* * `s` consist of only digits and English letters.
* 
*/
// @leetup=custom
// @leetup=code

import "strings"

func longestPalindrome(s string) string {
  sp := make([]rune, (len(s) * 2) + 1)
  
  j := 0
  for i, c := range s {
    sp[j + i] = '|'
    sp[j+ i + 1] = c
    j = j + 1
  }

  sp[len(s) * 2] = '|'

  maxP := string(s[0])
  // for every index, grow palindrome until failure
  for i := range sp {
    start := i
    end := i
    // grow from here
    j := 1
    for i - j >= 0 && i + j < len(sp) && sp[i-j] == sp[i+j] {
      start = i - j
      end = i + j
      j++
    }

    if end - start > len(maxP) {
      maxP = string(sp[start:end+1])
    }
  }

  longestP := strings.Builder{}
  for _, c := range maxP {
    if c != '|' {
      longestP.WriteRune(c)
    }
  }

  return longestP.String()
}

func _longestPalindrome(s string) string {
  n := len(s)
  curP := strings.Builder{} 
  curB := strings.Builder{}

  if n == 1 {
    return s
  }

  maxP := string(s[0])

  curLookbackIdx := 1


  for i, c := range s {
    if i - curLookbackIdx >= 0 && []rune(s)[i - curLookbackIdx] == rune(c) {
      curP.WriteRune(c)
      curB.WriteRune(c)
      curLookbackIdx += 2
      continue
    }

    // odd start case
    if i - (curLookbackIdx+1) >= 0 && curLookbackIdx == 1 && []rune(s)[i - (curLookbackIdx+1)] == rune(c) {
      curB.WriteRune([]rune(s)[i - 2])
      curP.WriteRune([]rune(s)[i - 1])
      curP.WriteRune(c)
      curLookbackIdx += 3
      continue
    }

    curB.WriteString(curP.String())

    if curB.Len() > len(maxP) {
      maxP = curB.String()
    }

    curLookbackIdx = 1
    curB.Reset()
    curP.Reset()
  }

  curB.WriteString(curP.String())

  if curB.Len() > len(maxP) {
    maxP = curB.String()
  }

  return maxP
}
// @leetup=code
