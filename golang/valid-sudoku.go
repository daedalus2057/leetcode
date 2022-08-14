package leetcode

// @leetup=custom
// @leetup=info id=36 lang=golang slug=valid-sudoku

/*
* Determine if a `9 x 9` Sudoku board is valid. Only the filled cells need to be
* validated according to the following rules:
* 
* 1. Each row must contain the digits `1-9` without repetition.
* 2. Each column must contain the digits `1-9` without repetition.
* 3. Each of the nine `3 x 3` sub-boxes of the grid must contain the digits `1-9`
*    without repetition.
* 
* Note:
* 
* * A Sudoku board (partially filled) could be valid but is not necessarily
*   solvable.
* * Only the filled cells need to be validated according to the mentioned rules.
* 
* 
* Example 1:
* 
* Input: board = 
* [["5","3",".",".","7",".",".",".","."]
* ,["6",".",".","1","9","5",".",".","."]
* ,[".","9","8",".",".",".",".","6","."]
* ,["8",".",".",".","6",".",".",".","3"]
* ,["4",".",".","8",".","3",".",".","1"]
* ,["7",".",".",".","2",".",".",".","6"]
* ,[".","6",".",".",".",".","2","8","."]
* ,[".",".",".","4","1","9",".",".","5"]
* ,[".",".",".",".","8",".",".","7","9"]]
* Output: true
* 
* Example 2:
* 
* Input: board = 
* [["8","3",".",".","7",".",".",".","."]
* ,["6",".",".","1","9","5",".",".","."]
* ,[".","9","8",".",".",".",".","6","."]
* ,["8",".",".",".","6",".",".",".","3"]
* ,["4",".",".","8",".","3",".",".","1"]
* ,["7",".",".",".","2",".",".",".","6"]
* ,[".","6",".",".",".",".","2","8","."]
* ,[".",".",".","4","1","9",".",".","5"]
* ,[".",".",".",".","8",".",".","7","9"]]
* Output: false
* Explanation: Same as Example 1, except with the 5 in the top left corner
*  being modified to 8. Since there are two 8's in the top left 3x3 sub-box, i
* t is invalid.
* 
* 
* Constraints:
* 
* * `board.length == 9`
* * `board[i].length == 9`
* * `board[i][j]` is a digit `1-9` or `'.'`.
* 
* 1. Each row must contain the digits `1-9` without repetition.
* 2. Each column must contain the digits `1-9` without repetition.
* 3. Each of the nine `3 x 3` sub-boxes of the grid must contain the digits `1-9`
*    without repetition.
*/
// @leetup=custom
// @leetup=code

import ( 
  "strconv"
  "unicode/utf8"
)


func isValidSudoku(board [][]byte) bool {

  
  // validate rows
  rowMaps := make([]map[int8]bool, 9)
  for i := range rowMaps {
    rowMaps[i] = make(map[int8]bool)
  }

  // validate columns
  colMaps := make([]map[int8]bool, 9)
  for i := range colMaps {
    colMaps[i] = make(map[int8]bool)
  }
  // validate 9x9
  cellMaps := make([]map[int8]bool, 9)

  for i := range cellMaps {
    cellMaps[i] = make(map[int8]bool)
  }

	getRune := func(b byte) rune {
		r, _ := utf8.DecodeRune([]byte{b})
		return r
	}

  maybeDigit := func(in byte) (int8, bool) {
    instr := string(getRune(in))

    if instr == "." {
      return 0, false
    }

    digit, err := strconv.ParseInt(instr, 10, 8)
    if err != nil {
      return 0, false
    }

    if digit < 1 || digit > 9 {
      return 0, false
    }

    return int8(digit), true
  }

  getCell := func(i int, j int) int {
    g := func(k int) int {
      if k < 3 {
        return 0
      }

      if k > 5 {
        return 2
      }

      return 1
    }

    return 3*g(i) + g(j)
  }

  for i := range board {
    for j, data := range board[i] {

      val, ok := maybeDigit(data)
      if !ok {
        continue
      }

      // we have a digit -- have we seen it?
      if rowMaps[j][val] { // we have seen this value before!
        return false
      }

      rowMaps[j][val] = true

      if colMaps[i][val] { // we have seen this value before!
        return false
      }

      colMaps[i][val] = true

      c := getCell(i, j) 
      
      if cellMaps[c][val] {
        return false
      }

      cellMaps[c][val] = true
    }
  }
    
  return true
}
// @leetup=code
