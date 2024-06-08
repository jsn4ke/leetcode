package leetcode_146_LRUCache

type node struct {
	prev, next *node

	key, val int
}

type LRUCache struct {
	capacity int
	head     *node
	hash     map[int]*node
}

func Constructor(capacity int) LRUCache {
	c := new(LRUCache)
	c.capacity = capacity
	return *c
}

func (l *LRUCache) add(n *node) {
	if nil == l.head {
		l.head = n
		n.prev = n
		n.next = n
		return
	}
	n.next = l.head
	n.prev = l.head.prev
	l.head.prev.next = n
	l.head.prev = n
	l.head = n
}

func (l *LRUCache) move2Head(n *node) {
	if n == l.head {
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = l.head
	n.prev = l.head.prev
	l.head.prev.next = n
	l.head.prev = n
	l.head = n
}

func (l *LRUCache) removeTail() {
	if nil == l.head {
		return
	}
	if l.head == l.head.next {
		delete(l.hash, l.head.key)
		l.head = nil
		return
	}
	tail := l.head.prev
	tail.prev.next = l.head
	l.head.prev = tail.prev
	delete(l.hash, tail.key)
}

func (l *LRUCache) Get(key int) int {
	v, ok := l.hash[key]
	if !ok {
		return -1
	}
	l.move2Head(v)
	return v.val
}

func (l *LRUCache) Put(key int, value int) {
	if v, ok := l.hash[key]; ok {
		v.val = value
		l.move2Head(v)
		return
	}
	if len(l.hash) >= l.capacity {
		l.removeTail()
	}
	n := &node{key: key, val: value}
	l.add(n)
	if nil == l.hash {
		l.hash = make(map[int]*node)
	}
	l.hash[key] = n
}
