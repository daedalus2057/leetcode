package algo

class SquaresOfSortedArraySpec extends munit.FunSuite {
  test("Squares should return a result") {
    val nums = Array(-4, -2, -1, 0, 3, 5)
    val res = SquaresOfSortedArray(nums)

    val expected = Array(0, 1, 4, 9, 16, 25)
    
    assertEquals(res.nums.toSeq, expected.toSeq)
  }

  test("Squares should be O(n)") {
    val nums = Array(-4, -2, -1, 0, 3, 5)
    val res = SquaresOfSortedArray(nums)

    val numsScaled = Array(-25, -12, -7, -4, -2, -1, 0, 3, 5, 6, 9, 13)
    val resScaled = SquaresOfSortedArray(numsScaled)

    assert((resScaled.ops / res.ops) <= 2)
  }


  test("All negative numbers should work") {
    val nums = Array(-25, -12, -7, -4, -2, -1)
    val res = SquaresOfSortedArray(nums)

    val expected = Array(1, 4, 16, 49, 144, 625)

    assertEquals(res.nums.toSeq, expected.toSeq)
  }

  test("All positive numbers should work") {
    val nums = Array(0, 1, 4, 9, 16, 25)
    val res = SquaresOfSortedArray(nums)

    val expected = Array(0, 1, 16, 81, 256, 625)

    assertEquals(res.nums.toSeq, expected.toSeq)
  }


  test("Should work for [-4,0,1,1]") {

    val nums = Array(-4,0,1,1)
    val res = SquaresOfSortedArray(nums)

    val expected = Array(0, 1, 1, 16)

    assertEquals(res.nums.toSeq, expected.toSeq)

  }

}
