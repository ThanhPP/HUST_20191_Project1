package main

type edge struct {
	node   string
	weight int
}

type graph struct {
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(start, des string, weight int) {
	g.nodes[start] = append(g.nodes[start], edge{node: des, weight: weight})
	g.nodes[des] = append(g.nodes[des], edge{node: start, weight: weight})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (g *graph) dijikstra(start, des string) (int, []string) {
	h := newHeap()
	h.push(path{
		cost:  0,
		nodes: []string{start},
	})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == des {
			return p.cost, p.nodes
		}

		for _, i := range g.getEdges(node) {
			if !visited[i.node] {
				h.push(path{
					cost:  p.cost + i.weight,
					nodes: append([]string{}, append(p.nodes, i.node)...),
				})
			}
		}

		visited[node] = true
	}

	return 0, nil
}
