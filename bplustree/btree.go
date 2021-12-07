package bplustree

type Item struct {
	key   int
	Value interface{}
}

// BTree
// Ref: https://www.cs.utexas.edu/users/djimenez/utsa/cs3343/lecture16.html
// degree d
//   - means the minimum chrildren a node can have, that means the min keys a node can have is: d - 1
//   - a node can have max 2*d - 1 keys
// btree properties
//   1. Every node has at most 2*d childrens
//   2. Root has at least 2 chrildren
//   3. A internal node with k chrildren, has k - 1 keys
//   4. All leaves appear in the same level
//   5. A node is full when it's number of keys is 2*d - 1
type BTree struct {
	degree int
	root   *Node
	length int
}

func (b *BTree) Find(key int) *Item {
	if b.root == nil {
		return nil
	}

	return b.root.find(key)
}

func (b *BTree) Insert(item *Item) *Item {
	if item == nil {
		return nil
	}

	if b.root == nil {
		b.root = new(Node)
		b.root.items = append(b.root.items, item)
		b.length++
		return nil
	}

	maxSize := b.maxSize()

	if len(b.root.items) >= maxSize {
		item1, second := b.root.split(maxSize / 2)
		root := b.root

		b.root = new(Node)
		b.root.items = append(b.root.items, item1)
		b.root.children = append(b.root.children, root, second)
	}

	itm := b.root.insert(item, maxSize)
	if itm == nil {
		b.length++
	}

	return itm
}

func (b *BTree) maxSize() int {
	return b.degree*2 - 1
}

func (b *BTree) Delete(key int) {
}

func (b *BTree) String() string {
	return ""
}

func (b *BTree) Iterate() []*Item {
	return nil
}
