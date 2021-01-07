package calcequation

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
