package leetcode_208_Trie

type node struct {
	char byte
	leaf bool
	next *level
}

type level struct {
	pair []*node
}

type Trie struct {
	root *level
}

func Constructor() Trie {
	t := new(Trie)
	return *t
}

func (t *Trie) Insert(word string) {
	if 0 == len(word) {
		return
	}
	var pre *node
	var root = t.root
	for i := range word {
		if nil == root {
			root = new(level)
			if nil == pre {
				t.root = root
			} else {
				pre.next = root
			}
		}
		c := word[i]
		var find *node
		for _, pair := range root.pair {
			if pair.char == c {
				find = pair
				break
			}
		}
		if nil == find {
			find = new(node)
			find.char = c
			root.pair = append(root.pair, find)
		}
		pre = find
		root = find.next
	}
	if nil != pre {
		pre.leaf = true
	}
}

func (t *Trie) Search(word string) bool {
	var last *node
	root := t.root
	for i := range word {
		if nil == root {
			return false
		}
		c := word[i]
		var find *node
		for _, pair := range root.pair {
			if pair.char == c {
				find = pair
				break
			}
		}
		if nil == find {
			return false
		}
		last = find
		root = find.next
	}
	return nil != last && last.leaf
}

func (t *Trie) StartsWith(prefix string) bool {
	root := t.root
	for i := range prefix {
		if nil == root {
			return false
		}
		c := prefix[i]
		var find *node
		for _, pair := range root.pair {
			if pair.char == c {
				find = pair
				break
			}
		}
		if nil == find {
			return false
		}
		root = find.next
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
