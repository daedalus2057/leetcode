package leetcode

import "testing"

const (
  DOT = byte(0x2e)
  ONE = byte(0x31)
  TWO = byte(0x32)
  THREE = byte(0x33)
  FOUR = byte(0x34)
  FIVE = byte(0x35)
  SIX = byte(0x36)
  SEVEN = byte(0x37)
  EIGHT = byte(0x38)
  NINE = byte(0x39)
)

func TestIsValidSudoku(t *testing.T) {
//  [["5","3",".",".","7",".",".",".","."]
//  ,["6",".",".","1","9","5",".",".","."]
//  ,[".","9","8",".",".",".",".","6","."]
//  ,["8",".",".",".","6",".",".",".","3"]
//  ,["4",".",".","8",".","3",".",".","1"]
//  ,["7",".",".",".","2",".",".",".","6"]
//  ,[".","6",".",".",".",".","2","8","."]
//  ,[".",".",".","4","1","9",".",".","5"]
//  ,[".",".",".",".","8",".",".","7","9"]]
  board := [][]byte{
    {FIVE, THREE, DOT, DOT, SEVEN, DOT, DOT, DOT, DOT},
    {SIX, DOT, DOT, ONE, NINE, FIVE, DOT, DOT, DOT},
    {DOT, NINE, EIGHT, DOT, DOT,DOT, DOT, SIX, DOT},
    {EIGHT, DOT, DOT, DOT, SIX, DOT, DOT, DOT, THREE},
    {FOUR, DOT, DOT, EIGHT, DOT, THREE, DOT, DOT, ONE},
    {SEVEN, DOT, DOT, DOT, TWO, DOT, DOT, DOT, SIX},
    {DOT, SIX, DOT, DOT, DOT, DOT, TWO, EIGHT, DOT},
    {DOT, DOT, DOT, FOUR, ONE, NINE, DOT, DOT, FIVE},
    {DOT, DOT, DOT, DOT, EIGHT, DOT, DOT, SEVEN, NINE},
  }

  if !isValidSudoku(board) {
    t.Fatal("Board was vaild but algo didn't think so")
  }

  board[0][0] = EIGHT

  if isValidSudoku(board) {
    t.Fatal("Board was invaild but algo didn't think so")
  }

}

func TestBoxValidity(t *testing.T) {
  /*
[[".",".",".",".","5",".",".","1","."]
,[".","4",".","3",".",".",".",".","."]
,[".",".",".",".",".","3",".",".","1"]
,["8",".",".",".",".",".",".","2","."]
,[".",".","2",".","7",".",".",".","."]
,[".","1","5",".",".",".",".",".","."]
,[".",".",".",".",".","2",".",".","."]
,[".","2",".","9",".",".",".",".","."]
,[".",".","4",".",".",".",".",".","."]]
*/

  board := [][]byte{
    {DOT, DOT, DOT, DOT, FIVE, DOT, DOT, ONE, DOT},
    {DOT, FOUR, DOT, THREE, DOT, DOT, DOT, DOT, DOT},
    {DOT, DOT, DOT, DOT, DOT, THREE, DOT, DOT, ONE},
    {EIGHT, DOT, DOT, DOT, DOT, DOT, DOT, TWO, DOT},
    {DOT, DOT, TWO, DOT, SEVEN, DOT, DOT, DOT, DOT},
    {DOT, ONE, FIVE, DOT, DOT, DOT, DOT, DOT, DOT},
    {DOT, DOT, DOT, DOT, DOT, TWO, DOT, DOT, DOT},
    {DOT, TWO, DOT, NINE, DOT, DOT, DOT, DOT, DOT},
    {DOT, DOT, FOUR, DOT, DOT, DOT, DOT, DOT, DOT},
  }

  
  if isValidSudoku(board) {
    t.Fatal("Board was invaild but algo didn't think so")
  }
}
