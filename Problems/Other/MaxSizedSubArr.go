package Problems

/*

Given an integer array nums and an integer k, return the maximum length of a subarray that sums to k. If there isn't one, return 0 instead.
Example 1:

Input: nums = [1,-1,5,-2,3], k = 3
Output: 4
Explanation: The subarray [1, -1, 5, -2] sums to 3 and is the longest.
Example 2:

Input: nums = [-2,-1,2,1], k = 1
Output: 2
Explanation: The subarray [-1, 2] sums to 1 and is the longest.
*/

// How does it work - indexing hash map

func maxSubArrayLen(nums []int, k int) int {
	prefixSum, longestSubarray := 0, 0
	indices := make(map[int]int, len(nums)+1)

	for i, v := range nums {
		prefixSum += v
		if prefixSum == k {
			longestSubarray = i + 1
		}

		if _, ok := indices[prefixSum-k]; ok {
			if longestSubarray < i-indices[prefixSum-k] {
				longestSubarray = i - indices[prefixSum-k]
			}
		}

		if _, ok := indices[prefixSum]; !ok {
			indices[prefixSum] = i
		}
	}
	return longestSubarray
}
