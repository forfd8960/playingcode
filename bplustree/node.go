package bplustree

import "sort"

type Items []*Item

type Node struct {
	items    Items
	children []*Node
}

// https://go.dev/play/p/5GLIFnMVaCY
func (items Items) find(key int) (int, bool) {
	i := sort.Search(len(items), func(i int) bool {
		return key < items[i].key
	})

	if i > 0 && items[i-1].key == key {
		return i - 1, true
	}

	return i, false
}

func (n *Node) find(key int) *Item {
	i, found := n.items.find(key)
	if found {
		return n.items[i]
	}

	if len(n.children) > 0 {
		return n.children[i].find(key)
	}

	return nil
}

func (n *Node) insert(item *Item) {
	if item == nil {
		return
	}
}

func (n *Node) split() {

}
