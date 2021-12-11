package bplustree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeInsert(t *testing.T) {
	cases := []struct {
		desc   string
		item   *Item
		n      *Node
		result *Node

		maxSize int
	}{
		{
			desc:    "insert to root",
			maxSize: 2,
			n:       &Node{},
			item:    &Item{key: 1},
			result: &Node{items: []*Item{
				{key: 1},
			}},
		},
		{
			desc:    "insert more items",
			maxSize: 2,
			item:    &Item{key: 2},
			n: &Node{items: []*Item{
				{key: 1},
			}},
			result: &Node{items: []*Item{
				{key: 1},
				{key: 2},
			}},
		},
		{
			desc:    "insert to children",
			maxSize: 2,
			item:    &Item{key: 5},
			n: &Node{
				items: []*Item{
					{key: 2},
				},
				children: []*Node{
					{
						items: []*Item{
							{key: 1},
						},
					},
					{
						items: []*Item{
							{key: 3},
						},
					},
				},
			},
			result: &Node{items: []*Item{
				{key: 2},
			},
				children: []*Node{
					{
						items: []*Item{
							{key: 1},
						},
					},
					{
						items: []*Item{
							{key: 3},
							{key: 5},
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.n.insert(tc.item, tc.maxSize)
			assert.Equal(t, tc.result, tc.n)
		})
	}
}
