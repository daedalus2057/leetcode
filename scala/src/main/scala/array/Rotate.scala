package array

import scala.collection.mutable.ArrayBuffer

import scala.annotation.tailrec

/*
 * Works but the ((i - k) % len) + len) % len was more or less
 * guessed at to make work. Getting a good understanding of this
 * problem is an excellent number theory challenge.
 */
def rotate(nums: Array[Int], k: Int): Unit = {
  Array.copy(
    nums.zipWithIndex.map( (v, i) => nums((((i - k) % nums.length) + nums.length) % nums.length) ), 0,
    nums, 0,
    nums.length)
}
