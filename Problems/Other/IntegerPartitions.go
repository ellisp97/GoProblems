package Problems

/*
An integer partition of n is a weakly decreasing list of positive integers which sum to n.

For example, there are 7 integer partitions of 5:

[5], [4,1], [3,2], [3,1,1], [2,2,1], [2,1,1,1], [1,1,1,1,1].

Write a function named partitions which returns the number of integer partitions of n. The function should be able to find the number of integer partitions of n for n as least as large as 100.
*/

// Use DynamicProgramming Logic from the coin change problem here
func Partitions(n int) int {
	// 1. First make array of n+1, and set 0 index to 1
	combinations := make([]int, n+1)
	// At index i, this is the total count of combinations
	// for the amount i
	combinations[0] = 1

	// 2. Loop through every coin (this instance n,n-1,...1)
	for coin := 1; coin < n+1; coin++ {
		// 3. For each coin loop through all amount
		for amount := 1; amount < len(combinations); amount++ {

			if amount >= coin {
				combinations[amount] += combinations[amount-coin]
			}
		}

	}
	return combinations[n]
}
