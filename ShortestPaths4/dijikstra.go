package main

func dijikstra(xStart, yStart, xDes, yDes int) (int, []nodeTemplate, [size][size]bool) {
	var startNode nodeTemplate
	startNode.x = xStart
	startNode.y = yStart
	h := newHeap()
	h.push(path{
		cost:  0,
		nodes: []nodeTemplate{startNode},
	})
	var visited [size][size]bool
	for len(*h.values) > 0 {
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node.x][node.y] {
			continue
		}

		if node.x == xDes && node.y == yDes {
			return p.cost, p.nodes, visited
		}

		for _, i := range getEdges(node.x, node.y) {
			if !visited[i.node.x][i.node.y] {
				tempNode := nodeTemplate{
					x: i.node.x,
					y: i.node.y,
				}
				h.push(path{
					cost:  p.cost + i.weight,
					nodes: append([]nodeTemplate{}, append(p.nodes, tempNode)...),
				})
			}
			visited[node.x][node.y] = true
		}
	}
	return 0, nil, visited
}
