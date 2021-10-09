package Problems

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	leftprefixarr, rightsuffixarr := make([]int, len(nums)), make([]int, len(nums))
	leftprefixarr[0], rightsuffixarr[len(nums)-1] = 1, 1

	for i := 1; i < len(nums); i++ {
		leftprefixarr[i] = nums[i-1] * leftprefixarr[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		rightsuffixarr[i] = nums[i+1] * rightsuffixarr[i+1]
	}

	for i := 0; i < len(nums); i++ {
		res[i] = leftprefixarr[i] * rightsuffixarr[i]
	}
	return res
}

// O(1) space complexity by reusing the res array
func productExceptSelf1(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1

	for i := 1; i < len(nums); i++ {
		res[i] = nums[i-1] * res[i-1]
	}

	num := 1
	for i := len(nums) - 2; i >= 0; i-- {
		num *= nums[i+1]
		res[i] *= num
	}
	return res
}
