package bplustree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTreeFind(t *testing.T) {
	cases := []struct {
		desc       string
		btree      *BTree
		key        int
		expectItem *Item
	}{
		{
			desc: "find in a empty tree",
			btree: &BTree{
				root: nil,
			},
			key:        1,
			expectItem: nil,
		},
		{
			desc: "item in the root",
			btree: &BTree{
				degree: 2,
				root: &Node{
					items: Items{
						{key: 1},
						{key: 5, Value: "5"},
						{key: 6},
					},
				},
				length: 3,
			},
			key:        5,
			expectItem: &Item{key: 5, Value: "5"},
		},
		{
			desc: "item in the child",
			btree: &BTree{
				degree: 2,
				root: &Node{
					items: Items{
						{key: 5},
					},
					children: []*Node{
						{
							items: Items{
								{key: 1},
							},
						},
						{
							items: Items{
								{key: 6},
								{key: 8, Value: "eight"},
							},
						},
					},
				},
				length: 4,
			},
			key:        8,
			expectItem: &Item{key: 8, Value: "eight"},
		},
		{
			desc: "item not in the tree",
			btree: &BTree{
				degree: 2,
				root: &Node{
					items: Items{
						{key: 5},
					},
					children: []*Node{
						{
							items: Items{
								{key: 1},
							},
						},
						{
							items: Items{
								{key: 6},
								{key: 8, Value: "eight"},
							},
						},
					},
				},
				length: 4,
			},
			key:        9,
			expectItem: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			result := tc.btree.Find(tc.key)
			assert.Equal(t, tc.expectItem, result)
		})
	}
}

func TestBtreeInsert(t *testing.T) {
	cases := []struct {
		desc        string
		btree       *BTree
		items       []*Item
		expectBtree *BTree
	}{
		{
			desc: "insert into root",
			btree: &BTree{
				degree: 3,
				root:   nil,
			},
			items: []*Item{
				{key: 1},
			},
			expectBtree: &BTree{
				degree: 3,
				root: &Node{
					items: Items{
						{key: 1},
					},
				},
				length: 1,
			},
		},
		{
			desc: "split root",
			btree: &BTree{
				degree: 2,
				root: &Node{
					items: Items{
						{key: 1},
						{key: 5},
						{key: 6},
					},
				},
				length: 3,
			},
			items: []*Item{
				{key: 8},
			},
			expectBtree: &BTree{
				degree: 2,
				root: &Node{
					items: Items{
						{key: 5},
					},
					children: []*Node{
						{
							items: Items{
								{key: 1},
							},
						},
						{
							items: Items{
								{key: 6},
								{key: 8},
							},
						},
					},
				},
				length: 4,
			},
		},
		{
			desc: "insert into child",
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			for _, itm := range tc.items {
				tc.btree.Insert(itm)
			}
			assert.Equal(t, tc.expectBtree, tc.btree)
		})
	}
}
