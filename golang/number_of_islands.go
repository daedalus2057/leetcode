package leetcode
// @leetup=custom
// @leetup=info id=200 lang=golang slug=number-of-islands

/*
* Given an `m x n` 2D binary grid `grid` which represents a map of `'1'`s (land)
* and `'0'`s (water), return *the number of islands*.
* 
* An island is surrounded by water and is formed by connecting adjacent lands
* horizontally or vertically. You may assume all four edges of the grid are all
* surrounded by water.
* 
* 
* Example 1:
* 
* Input: grid = [
*   ["1","1","1","1","0"],
*   ["1","1","0","1","0"],
*   ["1","1","0","0","0"],
*   ["0","0","0","0","0"]
* ]
* Output: 1
* 
* Example 2:
* 
* Input: grid = [
*   ["1","1","0","0","0"],
*   ["1","1","0","0","0"],
*   ["0","0","1","0","0"],
*   ["0","0","0","1","1"]
* ]
* Output: 3
* 
* 
* Constraints:
* 
* * `m == grid.length`
* * `n == grid[i].length`
* * `1 <= m, n <= 300`
* * `grid[i][j]` is `'0'` or `'1'`.
* 
*/
// @leetup=custom
// @leetup=code

func Union(a, b int, root_a []int) {
  rootA := Find(a, root_a)
  rootB := Find(b, root_a)

  if rootA != rootB {
    root_a[rootB] = rootA
  }
}

func Find(v int, root_a []int) (root int) {

  n := v
  p := root_a[v]

  for p != n {
    n = p
    p = root_a[n]
  }

  return p
}

func Init(root_a []int) {
  for i := range root_a {
    root_a[i] = i
  }
}

func Index(i, j, sz int) int {
  return sz*i + j
}

func numIslands(grid [][]byte) (int, []int) {
  M := len(grid)
  N := len(grid[0])

  root_a := make([]int, M * N)
  Init(root_a)

  i := 0
  for i < M { // Rows
    j := 0
    for j < N { // Cols
      // for each cell look right and down
      // if its a 1 then Union the two 
      if grid[i][j] == '1' {
        // Look Down
        if i < M - 1 {
          if grid[i + 1][j] == '1' {
            Union(Index(i, j, N), Index(i + 1, j, N), root_a)
          } 
        }

        // Look Right
        if j < N - 1 {
          if grid[i][j + 1] == '1' {
            Union(Index(i, j, N), Index(i, j+1, N), root_a)
          }
        }
      } else {
        // clear the root b/c this is not a node
        root_a[Index(i, j, N)] = -1
      }

      j = j + 1
    }

    i = i + 1
  }

  // now we look for all the roots and that is the number of islands
  num := 0
  for i, v := range root_a {
    if i == v {
      num = num + 1
    }
  }

  return num, root_a
}
// @leetup=code
