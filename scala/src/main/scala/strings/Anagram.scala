package strings

/*
 * Chars to be considered: alphanum
 * Chars to be submitted: ASCII printable chars
 * Target runtime: O(n)
 * Obtained runtime: O(n lg n)
 */
def isAnagramLogN(s: String, t: String): Boolean = {
  def normalize(s: String) = s.sorted.filter( _.isLetterOrDigit ).map(_.toLower)
  val a = normalize(s)
  val b = normalize(t)
  a == b
}

def inc(by: Int)(acc: Vector[Int], v: Int): Vector[Int] = 
  acc.updated(v, acc(v) + by)

def accumulate(f: (Vector[Int], Int) => Vector[Int]) = 
  (acc: Vector[Int], c: Char) =>
    if (c.isLetterOrDigit) f(acc, c.toUpper - 48) 
    else acc 
/**
 * Technically, this is not linear b/c Vector is not O(1)
 * on updates, but this maintains immutable execution
 * which in scala is considered a best practice. Furthermore
 * it would be trivial to change this to use an Array which
 * would make it truely linear
 */
def isAnagramLinear(s: String, t: String): Boolean =
  t.foldLeft(
    s.foldLeft(
      Vector.fill(43)(0)
    )(accumulate(inc(1)))
  )(accumulate(inc(-1))).filterNot( _ == 0 ) == Nil   


def isAnagram = isAnagramLinear _
