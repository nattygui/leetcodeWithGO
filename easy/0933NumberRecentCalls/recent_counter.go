package recentcounter

// RecentCounter 定义
type RecentCounter struct {
	records []int
}

// Constructor 构造RecentCounter
func Constructor() RecentCounter {
	return RecentCounter{
		records: make([]int, 0),
	}
}

// Ping 获取[t-3000, t] 以内的所有数
func (rc *RecentCounter) Ping(t int) int {
	count := 0

	min := t - 3000
	if min < 0 {
		min = 0
	}
	max := t
	for _, record := range rc.records {
		if record >= min && record <= max {
			count++
		}
	}
	rc.records = append(rc.records, t)
	return count + 1
}
