package strings

/*
 * Chars to be considered: alphanum
 * Chars to be submitted: ASCII printable chars
 * Target runtime: O(n)
 * Obtained runtime: O(n lg n)
 */
def isAnagram(s: String, t: String): Boolean = {
  def normalize(s: String) = s.sorted.filter( _.isLetterOrDigit ).map(_.toLower)
  val a = normalize(s)
  val b = normalize(t)
  a == b
}
