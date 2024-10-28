/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package main

// @lc code=start
// Time: O(n), space: O(n), n = number of nums
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for index, value := range nums {
		if pairIndex, ok := m[target-value]; ok {
			return []int{pairIndex, index}
		} else {
			m[value] = index
		}
	}
	panic("沒有合適的配對")
}

// @lc code=end
