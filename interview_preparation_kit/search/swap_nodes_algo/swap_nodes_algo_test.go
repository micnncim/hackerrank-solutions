package swap_nodes_algo

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_swapNodes(t *testing.T) {
	tests := []struct {
		name    string
		indexes [][]int32
		queries []int32
		want    [][]int32
	}{
		{
			indexes: [][]int32{
				{2, 3},
				{-1, -1},
				{-1, -1},
			},
			queries: []int32{1, 1},
			want: [][]int32{
				{3, 1, 2},
				{2, 1, 3},
			},
		},
		{
			indexes: [][]int32{
				{2, 3},
				{4, 5},
				{6, -1},
				{-1, 7},
				{8, 9},
				{10, 11},
				{12, 13},
				{-1, 14},
				{-1, -1},
				{15, -1},
				{16, 17},
				{-1, -1},
				{-1, -1},
				{-1, -1},
				{-1, -1},
				{-1, -1},
				{-1, -1},
			},
			queries: []int32{2, 3},
			want: [][]int32{
				{14, 8, 5, 9, 2, 4, 13, 7, 12, 1, 3, 10, 15, 6, 17, 11, 16},
				{9, 5, 14, 8, 2, 13, 7, 12, 4, 1, 3, 17, 11, 16, 6, 10, 15},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNodes(tt.indexes, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	tests := []struct {
		name string
		n    *node
		k    int32
		want *node
	}{
		{
			n: &node{
				index: 1,
				left:  &node{index: 2},
				right: &node{index: 3},
			},
			k: 1,
			want: &node{
				index: 1,
				left:  &node{index: 3},
				right: &node{index: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.n, tt.k)
			assert.Equal(t, tt.want, tt.n)
		})
	}
}

func Test_binaryTree(t *testing.T) {
	tests := []struct {
		name    string
		indexes [][]int32
		want    []*node
	}{
		{
			indexes: [][]int32{
				{2, 3},
				{-1, -1},
				{-1, -1},
			},
			want: []*node{
				{
					index: 1,
					depth: 1,
					left: &node{
						index: 2,
						depth: 2,
					},
					right: &node{
						index: 3,
						depth: 2,
					},
				},
				{
					index: 2,
					depth: 2,
				},
				{
					index: 3,
					depth: 2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binaryTree(tt.indexes)
			assert.Equal(t, tt.want, got)
		})
	}
}
