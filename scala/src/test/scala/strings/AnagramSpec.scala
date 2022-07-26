package strings

import munit._
import org.scalacheck.Prop._
import org.scalacheck.Gen

class AnagramPropSpec extends ScalaCheckSuite {

  val asciiGen = Gen.containerOf[Array, Char](Gen.choose[Char](0x20,0x7e)).map(_.mkString)
  val rando = scala.util.Random()
  def randomBool = rando.nextBoolean

  property("check equivelance of ascii alpha numeric with whitespace and punctuation anagrams") {
    forAll(asciiGen) {
      (s: String) =>
        // we should have a random string with some non-alpha chars 
        // now, lets filter out some of the non-alpha chars and shuffle
        val t = rando.shuffle(s.filter( _.isLetterOrDigit || randomBool )).toString
        assert(isAnagram(s, t))
        

    }

  }
}

class AnagramSpec extends munit.FunSuite {

  test("Empty string should be an anagram of itself") {
    assert(isAnagram("", ""))
  }

}
