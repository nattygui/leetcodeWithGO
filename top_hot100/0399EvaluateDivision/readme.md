# 399. 除法求值
给出方程式 A / B = k, 其中 A 和 B 均为用字符串表示的变量， k 是一个浮点型数字。根据已知方程式求解问题，并返回计算结果。如果结果不存在，则返回 -1.0。

输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

示例 1：
```
输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
解释：
给定：a / b = 2.0, b / c = 3.0
问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
返回：[6.0, 0.5, -1.0, 1.0, -1.0 ]
```
示例 2：
```
输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
输出：[3.75000,0.40000,5.00000,0.20000]
```

## 解法

```go
type node struct {
	nv    []float64
	next  []*node
	pv    []float64
	prev  []*node
	visit bool
	value float64
	group int
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// -----
	nodeCache := make(map[string]*node)
	nodes := make([]*node, 0)
	for index, e := range equations {
		x, y := e[0], e[1]
		v1, ok1 := nodeCache[x]
		v2, ok2 := nodeCache[y]
		if !ok1 && !ok2 {
			node1 := &node{}
			node2 := &node{}
			node1.next = append(node1.next, node2)
			node1.nv = append(node1.nv, values[index])
			node2.prev = append(node2.prev, node1)
			node2.pv = append(node2.pv, 1.0/values[index])
			nodes = append(nodes, node1, node2)
			nodeCache[x] = node1
			nodeCache[y] = node2
		} else if ok1 && !ok2 {
			node2 := &node{}
			v1.next = append(v1.next, node2)
			v1.nv = append(v1.nv, values[index])
			node2.prev = append(node2.prev, v1)
			node2.pv = append(node2.pv, 1.0/values[index])
			nodes = append(nodes, node2)
			nodeCache[y] = node2
		} else if ok2 && !ok1 {
			node1 := &node{}
			v2.next = append(v2.next, node1)
			v2.nv = append(v2.nv, 1.0/values[index])
			node1.prev = append(node1.prev, v2)
			node1.pv = append(node1.pv, values[index])
			nodes = append(nodes, node1)
			nodeCache[x] = node1
		} else {
			v1.next = append(v1.next, v2)
			v1.nv = append(v1.nv, values[index])

			v2.prev = append(v2.prev, v1)
			v2.pv = append(v2.pv, 1.0/values[index])
		}
	}
	// 每个节点的值
	group := 1
	for _, n := range nodes {
		if !n.visit {
			n.value = 1.0
			n.group = group
			initNodes(n)
		}
		group++
	}

	result := make([]float64, 0, len(queries))
	for _, q := range queries {
		x, y := q[0], q[1]
		v1, ok1 := nodeCache[x]
		v2, ok2 := nodeCache[y]
		if ok1 && ok2 && v1.group == v2.group {
			result = append(result, v1.value/v2.value)
		} else {
			result = append(result, -1)
		}
	}
	return result
}

func initNodes(n *node) {
	if n.visit {
		return
	}
	n.visit = true
	for index, subNode := range n.next {
		if subNode.visit {
			continue
		}
		subNode.value = n.value / n.nv[index]
		subNode.group = n.group
	}
	for index, subNode := range n.prev {
		if subNode.visit {
			continue
		}
		subNode.value = n.value / n.pv[index]
		subNode.group = n.group
	}
	for _, subNode := range n.next {
		initNodes(subNode)
	}
	for _, subNode := range n.prev {
		initNodes(subNode)
	}
}
```