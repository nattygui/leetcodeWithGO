package leastinterval

import "math"

type node struct {
	task     byte
	number   int
	interval int
}

func leastInterval(tasks []byte, n int) int {
	taskMap := make(map[byte]*node, 0)
	for _, task := range tasks {
		if n, ok := taskMap[task]; ok {
			n.number++
		} else {
			taskMap[task] = &node{
				task:   task,
				number: 1,
			}
		}
	}
	runned := 0
	interval := 0
	for {
		if runned == len(tasks) {
			break
		}
		max := math.MinInt64
		curtask := byte('0')
		for k, v := range taskMap {
			if v.interval > 0 {
				v.interval--
				continue
			}
			if max < v.number {
				max = v.number
				curtask = k
			}
		}
		if curtask == '0' {
			interval++
			continue
		}
		taskMap[curtask].number--
		if taskMap[curtask].number == 0 {
			delete(taskMap, curtask)
		} else {
			taskMap[curtask].interval = n
		}
		runned++
		interval++
	}
	return interval
}
