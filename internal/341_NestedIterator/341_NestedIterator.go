package leetcode_341_NestedIterator

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return true
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 1
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return nil }

// 多级栈
type NestedIterator struct {
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{stack: [][]*NestedInteger{nestedList}}
}

func (this *NestedIterator) Next() int {
	top := this.stack[len(this.stack)-1]
	cur := top[0]
	top = top[1:]
	this.stack[len(this.stack)-1] = top
	return cur.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		top := this.stack[len(this.stack)-1]
		if len(top) == 0 {
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}
		cur := top[0]
		top = top[1:]
		if cur.IsInteger() {
			return true
		}
		this.stack = append(this.stack, cur.GetList())
	}
	return false
}

func traverse(in []*NestedInteger) []int {
	res := make([]int, 0)
	for _, v := range in {
		if v.IsInteger() {
			res = append(res, v.GetInteger())
		} else {
			res = append(res, traverse(v.GetList())...)
		}
	}
	return res
}
