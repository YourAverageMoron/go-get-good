package main

func ProductOfArrayExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
    res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
    postfix := 1
    for i := len(nums) -2; i >=0; i-- {
        postfix *= nums[i+1]
        res[i] *= postfix
    }
	return res
}
