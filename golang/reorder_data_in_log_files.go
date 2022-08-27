package leetcode
// @leetup=custom
// @leetup=info id=937 lang=golang slug=reorder-data-in-log-files

/*
* You are given an array of `logs`. Each log is a space-delimited string of words,
* where the first word is the identifier.
* 
* There are two types of logs:
* 
* * Letter-logs: All words (except the identifier) consist of lowercase English
*   letters.
* * Digit-logs: All words (except the identifier) consist of digits.
* 
* Reorder these logs so that:
* 
* 1. The letter-logs come before all digit-logs.
* 2. The letter-logs are sorted lexicographically by their contents. If their
*    contents are the same, then sort them lexicographically by their identifiers.
* 3. The digit-logs maintain their relative ordering.
* 
* Return *the final order of the logs*.
* 
* 
* Example 1:
* 
* Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","
* let3 art zero"]
* Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","d
* ig2 3 6"]
* Explanation:
* The letter-log contents are all different, so their ordering is "art can", "art 
* zero", "own kit dig".
* The digit-logs have a relative order of "dig1 8 1 5 1", "dig2 3 6".
* 
* Example 2:
* 
* Input: logs = ["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act
*  zoo"]
* Output: ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"]
* 
* 
* Constraints:
* 
* * `1 <= logs.length <= 100`
* * `3 <= logs[i].length <= 100`
* * All the tokens of `logs[i]` are separated by a single space.
* * `logs[i]` is guaranteed to have an identifier and at least one word after the
*   identifier.
* 
*/
// @leetup=custom
// @leetup=code

import ( 
  "sort"
  "strings"
  "strconv"
)

func logIsText(log string) bool {
  strs := strings.Split(log, " ")
  _, err := strconv.Atoi(strs[1])
  if err != nil {
    return true
  }

  return false
}

func extractIdentAndRaw(log string) (string, []string) {
  raw := strings.Split(log, " ")
  ident := raw[0]
  return ident, raw[1:]
}

// mutating
func sortStringLogs(logLines []string) []string {
  // flip, sort, unflip
  sorted := make([]string, len(logLines))
  for i, line := range logLines {
    ident, raw := extractIdentAndRaw(line)
    raw = append(raw, ident)
    sorted[i] = strings.Join(raw, " ")
  }

  sort.Strings(sorted)

  // unflip
  for i, line := range sorted {
    raw := strings.Split(line, " ")
    ident := raw[len(raw)-1]
    raw = raw[:len(raw)-1]
    sorted[i] = strings.Join(append([]string{ ident }, raw...), " ")
  }

  return sorted
}

func arrayToLogLine(strA []string) string {
  return strings.Join(strA, " ")
}

func logSplitStr(log string) (string, []string) {
  return extractIdentAndRaw(log)
}

func reorderLogFiles(input []string) []string {
  nums := make([]string, 0)
  logs := make([]string, 0)
  for _, log := range input {
    if !logIsText(log) {
      nums = append(nums, log)
      continue
    }

    logs = append(logs, log)

  }
  logs = sortStringLogs(logs)
  return append(logs, nums...) 
}
// @leetup=code
