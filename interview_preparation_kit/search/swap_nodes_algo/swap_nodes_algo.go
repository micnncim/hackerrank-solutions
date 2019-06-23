package swap_nodes_algo

type node struct {
	index       int32
	depth       int32
	left, right *node
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	results := make([][]int32, 0, len(queries))
	tree := binaryTree(indexes)
	for _, k := range queries {
		swap(tree[0], k)
		result := make([]int32, 0, len(tree))
		traverseInOrder(tree[0], func(i interface{}) {
			result = append(result, i.(int32))
		})
		results = append(results, result)
	}
	return results
}

func binaryTree(indexes [][]int32) []*node {
	nodes := make([]*node, len(indexes))
	for i := range nodes {
		nodes[i] = &node{}
	}
	nodes[0] = &node{index: 1}
	for i, v := range indexes {
		nodes[i].depth = int32(i/2 + 1)
		if v[0] != -1 {
			nodes[i].left = &node{index: v[0]}
		}
		if v[1] != -1 {
			nodes[i].right = &node{index: v[1]}
		}
	}
	return nodes
}

func swap(n *node, k int32) {
	if n == nil {
		return
	}
	if n.depth%k == 0 {
		n.left, n.right = n.right, n.left
	}
	swap(n.left, k)
	swap(n.right, k)
}

func traverseInOrder(n *node, f func(interface{})) {
	if n == nil {
		return
	}
	traverseInOrder(n.left, f)
	f(n.index)
	traverseInOrder(n.right, f)
}
