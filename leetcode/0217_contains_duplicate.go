package main

func ContainsDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _, item := range nums {
        _, exists := m[item]
        if exists {
            return true
        }
        m[item] = true
    }
    return false
}
