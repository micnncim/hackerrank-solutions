package roads_and_libraries

import (
	"sync"
)

func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	clib := int64(c_lib)
	croad := int64(c_road)

	graph := make([]*node, n)
	for i := range graph {
		graph[i] = &node{neighbors: []int64{}}
	}
	for _, city := range cities {
		graph[city[0]-1].neighbors = append(graph[city[0]-1].neighbors, int64(city[1]-1))
		graph[city[1]-1].neighbors = append(graph[city[1]-1].neighbors, int64(city[0]-1))
	}
	var cost int64
	q := newQueue()
	for i := 0; i < len(graph); i++ {
		if graph[i].visited {
			continue
		}
		graph[i].visited = true
		cost += clib
		q.push(graph[i])
		for !q.isEmpty() {
			node := q.pop()
			for _, neighbor := range node.neighbors {
				if !graph[neighbor].visited {
					graph[neighbor].visited = true
					cost += min(clib, croad)
					q.push(graph[neighbor])
				}
			}
		}
	}
	return cost
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

type node struct {
	neighbors []int64
	visited   bool
}

type queue struct {
	values []*node
	size   int
	mu     *sync.Mutex
}

func newQueue() *queue {
	return &queue{
		values: make([]*node, 0),
		mu:     new(sync.Mutex),
	}
}

func (q *queue) isEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.size == 0
}

func (q *queue) pop() *node {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.size == 0 {
		return nil
	}
	v := q.values[0]
	if q.size == 1 {
		q.values = make([]*node, 0)
	} else {
		q.values = q.values[1:]
	}
	q.size--
	return v
}

func (q *queue) push(v *node) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.values = append(q.values, v)
	q.size++
}
