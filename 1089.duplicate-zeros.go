/*
 * @lc app=leetcode id=1089 lang=golang
 *
 * [1089] Duplicate Zeros
 */
package main

// @lc code=start
// Time: O(n), space: O(n), n = length of arr
type queue struct {
	container []interface{}
}

func (q *queue) Enqueue(data interface{}) {
	q.container = append(q.container, data)
}

func (q *queue) Dequeue() interface{} {
	if q.IsEmpty() == true {
		panic("queue is empty!")
	}
	data := q.container[0]
	q.container = q.container[1:]
	return data
}

func (q *queue) IsEmpty() bool {
	if len(q.container) == 0 {
		return true
	}
	return false
}

func duplicateZeros(arr []int) {
	q := queue{}
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
