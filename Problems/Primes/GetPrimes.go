package Problems

/*
Get all the prime numbers up to n

Idea: To do this implement Sieve of Eratosthenes

It does so by iteratively marking as composite the multiples
of each prime
*/

func SieveOfEratosthenes(n int) []bool {
	if n == 0 {
		return []bool{}
	}
	arr := make([]bool, n+1)

	//Initialise the array with true prime values
	for i, _ := range arr {
		arr[i] = true
	}
	//Mark 0 and 1 as not prime trivially
	arr[0], arr[1] = false, false

	for p := 2; p*p <= n; p++ {
		if arr[p] {
			// Update all multiples of p, marking them false
			// as they're divisible to 0
			for i := p * 2; i <= n; i += p {
				arr[i] = false
			}
		}
	}
	return arr
}
