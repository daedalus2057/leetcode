package algo

import scala.annotation.tailrec

case class SquaresOfSortedArrayResult(nums: Array[Int], ops: Int)
def SquaresOfSortedArray(nums: Array[Int]): SquaresOfSortedArrayResult = {
  var ops = 0

  def sq(i: Int) = i * i

  def swap(left: Int, right: Int) = {
    val tmp = nums( left )
    nums( left ) = nums( right )
    nums( right ) = tmp
  }

  @tailrec
  def reverse(left: Int, right: Int): Unit = {
    if (left <= right) {
      swap(left, right)
      reverse(left + 1, right - 1)
    }
  }

  @tailrec
  def go(left: Int, right: Int, acc: List[Int]): List[Int] = {
    if left > right then return acc

    if nums( left ).abs > nums( right ).abs then
      go(left + 1, right, sq(nums( left )) :: acc)
    else
      go(left, right - 1, sq(nums( right )) :: acc)
    end if
  }

  val squares = go(0, nums.size - 1, List[Int]())

  ops = ops + nums.size
  SquaresOfSortedArrayResult(squares.toArray, ops)
}
