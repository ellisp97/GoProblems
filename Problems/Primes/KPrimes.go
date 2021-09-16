package Problems

/*
Task:
Complete the function count_Kprimes (or countKprimes, count-K-primes, kPrimes) which is given parameters k, start, end (or nd) and returns an array (or a list or a string depending on the language - see "Solution" and "Sample Tests") of the k-primes between start (inclusive) and end (inclusive).

Example:
countKprimes(5, 500, 600) --> [500, 520, 552, 567, 588, 592, 594]
*/
func PrimeFactors(n int) int {
	res, fac := []int{}, 2
	for fac*fac <= n {
		for n%fac == 0 {
			n /= fac
			res = append(res, fac)
		}
		fac++
	}
	if n > 1 {
		res = append(res, n)
	}
	return len(res)
}

func CountKPrimes(k, start, end int) []int {
	res := []int{}
	for i := start; i <= end; i++ {
		if PrimeFactors(i) == k {
			res = append(res, i)
		}
	}
	return res
}
