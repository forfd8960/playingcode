package bplustree

type Item struct {
	key   int64
	Value interface{}
}

type Items []*Item

type Node struct {
	items    Items
	children Items
}

type BTree struct {
	degree int
	root   *Node
}

func (b *BTree) Search(key int64) *Item {
	return nil
}

func (b *BTree) Insert(item *Item) {

}

func (b *BTree) Delete(key int64) {

}

func (b *BTree) String() string {
	return ""
}

func (b *BTree) Iterate() []*Item {
	return nil
}
