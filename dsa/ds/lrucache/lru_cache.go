package lrucache

/* @tags: doubly_linked_list,lru,map,cache,linked_list */

type LRUCache struct {
	Head *Node
	Tail *Node
	HT   map[int]*Node
	Cap  int
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{HT: make(map[int]*Node, capacity+1), Cap: capacity}
}

func (lc *LRUCache) Get(key int) int {
	node, ok := lc.HT[key]
	if ok {
		lc.Remove(node)
		lc.Add(node)

		return node.Val
	}

	return -1
}

func (lc *LRUCache) Put(key, value int) {
	node, ok := lc.HT[key]
	if ok {
		node.Val = value
		lc.Remove(node)
		lc.Add(node)

		return
	} else {
		node = &Node{Key: key, Val: value}
		lc.HT[key] = node
		lc.Add(node)
	}

	if len(lc.HT) > lc.Cap {
		delete(lc.HT, lc.Tail.Key)
		lc.Remove(lc.Tail)
	}
}

func (lc *LRUCache) Add(node *Node) {
	node.Prev = nil

	node.Next = lc.Head
	if lc.Head != nil {
		lc.Head.Prev = node
	}

	lc.Head = node
	if lc.Tail == nil {
		lc.Tail = node
	}
}

func (lc *LRUCache) Remove(node *Node) {
	if node != lc.Head {
		node.Prev.Next = node.Next
	} else {
		lc.Head = node.Next
	}

	if node != lc.Tail {
		node.Next.Prev = node.Prev
	} else {
		lc.Tail = node.Prev
	}
}
