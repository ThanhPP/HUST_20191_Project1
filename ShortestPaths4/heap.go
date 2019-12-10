package main

import hp "container/heap"

type path struct {
	cost  int
	nodes []nodeTemplate
}

type pathA struct {
	heuristic float64
	cost      int
	nodes     []nodeTemplate
}

type minPath []path
type minPathA []pathA

func (h minPath) Len() int {
	return len(h)
}
func (h minPathA) Len() int {
	return len(h)
}

func (h minPath) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}
func (h minPathA) Less(i, j int) bool {
	return h[i].heuristic < h[j].heuristic
}

func (h minPath) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minPathA) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}
func (h *minPathA) Push(x interface{}) {
	*h = append(*h, x.(pathA))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *minPathA) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type heap struct {
	values *minPath
}
type heapA struct {
	values *minPathA
}

func newHeap() *heap {
	return &heap{values: &minPath{}}
}
func newHeapA() *heapA {
	return &heapA{values: &minPathA{}}
}

func (h *heap) push(p path) {
	hp.Push(h.values, p)
}
func (h *heapA) push(p pathA) {
	hp.Push(h.values, p)
}

func (h *heap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}
func (h *heapA) pop() pathA {
	i := hp.Pop(h.values)
	return i.(pathA)
}
