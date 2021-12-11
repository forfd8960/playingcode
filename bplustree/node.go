package bplustree

import "sort"

type Items []*Item

type Node struct {
	items    Items
	children []*Node
}

func insertAt(nodes []*Node, i int, n *Node) []*Node {
	nodes = append(nodes, nil)
	if i < len(nodes) {
		copy(nodes[i+1:], nodes[i:])
	}
	nodes[i] = n
	return nodes
}

// https://go.dev/play/p/5GLIFnMVaCY
func (items *Items) find(key int) (int, bool) {
	i := sort.Search(len(*items), func(i int) bool {
		return key < (*items)[i].key
	})

	if i > 0 && (*items)[i-1].key == key {
		return i - 1, true
	}

	return i, false
}

func (items *Items) insertAt(i int, item *Item) {
	*items = append(*items, nil)
	if i < len(*items) {
		copy((*items)[i+1:], (*items)[i:])
	}

	(*items)[i] = item
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

// https://webdocs.cs.ualberta.ca/~holte/T26/ins-b-tree.html
func (n *Node) insert(item *Item, maxSize int) *Item {
	i, found := n.items.find(item.key)
	if found {
		out := n.items[i]
		n.items[i] = item
		return out
	}

	if len(n.children) <= 0 {
		n.items.insertAt(i, item)
		return nil
	}

	if n.needSplitChildren(i, maxSize) {
		currentI := n.items[i]
		switch {
		case currentI.key < item.key: // insert item large than current key, insert it in right sub tree
			i++
		case currentI.key > item.key: // insert item less than current key, insert it in left sub tree
		default:
			n.items[i] = item
			return currentI
		}
	}

	return n.children[i].insert(item, maxSize)
}

func (n *Node) needSplitChildren(i int, maxSize int) bool {
	if len(n.children[i].items) < maxSize {
		return false
	}

	chd := n.children[i]
	item, second := chd.split(maxSize / 2)

	n.items.insertAt(i, item)
	n.children = insertAt(n.children, i+1, second)
	return true
}

func (n *Node) split(i int) (*Item, *Node) {
	item := n.items[i]
	second := new(Node)
	second.items = append(second.items, n.items[i+1:]...)
	n.items = n.items[:i]
	if len(n.children) > 0 {
		second.children = append(second.children, n.children[i+1:]...)
		n.children = n.children[:i+1]
	}

	return item, second
}
