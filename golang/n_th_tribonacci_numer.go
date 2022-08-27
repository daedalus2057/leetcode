package leetcode
// @leetup=custom
// @leetup=info id=1137 lang=golang slug=n-th-tribonacci-number

/*
* The Tribonacci sequence Tn is defined as follows:
* 
* T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
* 
* Given `n`, return the value of Tn.
* 
* 
* Example 1:
* 
* Input: n = 4
* Output: 4
* Explanation:
* T_3 = 0 + 1 + 1 = 2
* T_4 = 1 + 1 + 2 = 4
* 
* Example 2:
* 
* Input: n = 25
* Output: 1389537
* 
* 
* Constraints:
* 
* * `0 <= n <= 37`
* * The answer is guaranteed to fit within a 32-bit integer, ie. `answer <= 2^31 -
*   1`.
* 
*/
// @leetup=custom
// @leetup=code

func tribonacci(n int) int {
  tris := make([]int, n + 1)
  tris[0] = 0
  if n > 0 {
    tris[1] = 1
  }

  if n > 1 {
    tris[2] = 1
  }

  i := 3
  for i < n+1 {
    tris[i] = tris[i-1] + tris[i - 2] + tris[i - 3] 

    i = i + 1
  }
    
  return tris[n]
}
// @leetup=code
