/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    lo, hi := 0, n
	for hi > lo + 1 {
		mid := lo + (hi - lo)/2
		response := guess(mid)
		if response == 0 {
			return mid
		} else if response == -1 {
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}
