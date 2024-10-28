/*
 * @lc app=leetcode id=1089 lang=golang
 *
 * [1089] Duplicate Zeros
 */
package main

// @lc code=start
// Time: O(n), space: O(n), n = length of arr
type queue struct {
	container []interface
}

func (queue) Enqueue(data interface) {
	
}



func duplicateZeros(arr []int) {
	q := queue.New()
	for index, value := range arr {
		if value == 0 {
			q.Enqueue(0)
		}
		q.Enqueue(value)
		front := q.Dequeue().(int)
		arr[index] = front
	}
}

// @lc code=end
