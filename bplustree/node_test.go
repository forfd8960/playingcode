package bplustree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveItem(t *testing.T) {
	items := []*Item{
		{key: 1},
		{key: 3},
		{key: 6},
		{key: 8},
	}

	itms := (Items)(items)
	itms.deleteAt(2)

	expect := []*Item{
		{key: 1},
		{key: 3},
		{key: 8},
	}
	assert.Equal(t, (Items)(expect), itms)

	items = []*Item{
		{key: 1},
		{key: 3},
		{key: 6},
		{key: 8},
	}
	itms = (Items)(items)
	itms.deleteAt(3)
	expect = []*Item{
		{key: 1},
		{key: 3},
		{key: 6},
	}
	assert.Equal(t, (Items)(expect), itms)
}

func TestNodeFind(t *testing.T) {
	cases := []struct {
		desc   string
		n      *Node
		key    int
		expect *Item
	}{
		{
			desc: "key in node",
			n: &Node{
				items: Items{
					{key: 1},
					{key: 8},
					{key: 9},
				},
			},
			key:    9,
			expect: &Item{key: 9},
		},
		{
			desc: "key is in child",
			n: &Node{
				items: Items{
					{key: 2},
				},
				children: []*Node{
					{
						items: Items{
							{key: 1, Value: "1"},
						},
					},
					{
						items: Items{
							{key: 5},
						},
					},
				},
			},
			key:    5,
			expect: &Item{key: 5},
		},
		{
			desc: "key is not in node",
			n: &Node{
				items: Items{
					{key: 1},
					{key: 8},
					{key: 9},
				},
			},
			key:    2,
			expect: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			item := tc.n.find(tc.key)
			assert.Equal(t, tc.expect, item)
		})
	}
}

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
			maxSize: 3,
			n:       &Node{},
			item:    &Item{key: 1},
			result: &Node{items: []*Item{
				{key: 1},
			}},
		},
		{
			desc:    "insert more items",
			maxSize: 3,
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
			maxSize: 3,
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
		{
			desc:    "insert to children and split",
			maxSize: 3,
			item:    &Item{key: 8},
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
							{key: 5},
							{key: 6},
						},
					},
				},
			},
			result: &Node{
				items: []*Item{
					{key: 2},
					{key: 5},
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
					{
						items: []*Item{
							{key: 6},
							{key: 8},
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
