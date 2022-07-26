package array

class RotateSpec extends munit.FunSuite {
  test("Rotate should work with len 4 and k 2") {
    var test = Array[Int](-1, -100, 3, 99)
    val k = 2

    var expect = Array[Int](3, 99, -1, -100)
    rotate(test, k)
    assertEquals(test.toSeq, expect.toSeq)
  }

}
