package leetcode
// @leetup=custom
// @leetup=info id=819 lang=golang slug=most-common-word

/*
* Given a string `paragraph` and a string array of the banned words `banned`,
* return *the most frequent word that is not banned*. It is guaranteed there
* is at least one word that is not banned, and that the answer is unique.
* 
* The words in `paragraph` are case-insensitive and the answer should be
* returned in lowercase.
* 
* 
* Example 1:
* 
* Input: paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
* , banned = ["hit"]
* Output: "ball"
* Explanation: 
* "hit" occurs 3 times, but it is a banned word.
* "ball" occurs twice (and no other word does), so it is the most frequent non-ban
* ned word in the paragraph. 
* Note that words in the paragraph are not case sensitive,
* that punctuation is ignored (even if adjacent to words, such as "ball,"), 
* and that "hit" isn't the answer even though it occurs more because it is banned.
* 
* Example 2:
* 
* Input: paragraph = "a.", banned = []
* Output: "a"
* 
* 
* Constraints:
* 
* * `1 <= paragraph.length <= 1000`
* * paragraph consists of English letters, space `' '`, or one of the symbols:
*   `"!?',;."`.
* * `0 <= banned.length <= 100`
* * `1 <= banned[i].length <= 10`
* * `banned[i]` consists of only lowercase English letters.
* 
*/
// @leetup=custom
// @leetup=code

import "strings"

func getWords(p string, b []string) []string {
  builder := strings.Builder{}
  words := make([]string, 0)

  bannedMap := make(map[string]bool)
  for _, s := range b {
    bannedMap[strings.ToLower(s)] = true
  }
  
  splitters := make(map[rune]bool)
  puncs := []rune{ ' ', '!', '?', '\'', ';', ',', '.', '"', '`' } 
  for _, r := range puncs {
    splitters[r] = true  
  }

  for _, c := range p {
    if _, ok := splitters[c]; ok {
      if len(builder.String()) > 0 {
        words = append(words, builder.String())
        builder.Reset()
      }
      continue
    }

    builder.WriteRune(c)
  }

  if len(builder.String()) > 0 {
    words = append(words, builder.String())
  }

  return words
}

func mostCommonWord(paragraph string, banned []string) string {

  trimPunc := func(word string) string {
    puncs := []rune{ '!', '?', '\'', ',', '.', '"', '`' }
    for _, p := range puncs {
      if strings.HasPrefix(word, string(p)) {
        word = word[1:]
      }

      if strings.HasSuffix(word, string(p)) {
        word = word[:len(word)-1]
      }
    }

    return word
  }

  bannedMap := make(map[string]bool)

  for _, b := range banned {
    bannedMap[trimPunc(strings.ToLower(b))] = true
  }

  words := getWords(paragraph, banned) // strings.Split(paragraph, " ")
  freqMap := make(map[string]int)

  for _, word := range words {
    lowerWord := trimPunc(strings.ToLower(word))
    if _, x := bannedMap[lowerWord]; x {
      continue
    }
    
    if _, x := freqMap[lowerWord]; x {
      freqMap[lowerWord] = freqMap[lowerWord] + 1
      continue
    }

    freqMap[lowerWord] = 1
  }

  maxCount := -1
  maxWord := ""

  for w, c := range freqMap {
    if c > maxCount {
      maxCount = c
      maxWord = w
    }
  }

  return maxWord
}
// @leetup=code
