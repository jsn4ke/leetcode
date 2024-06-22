package leetcode_716_MaxStack

type MaxStack struct {
	stack  []int
	maxIdx []int
}

func Constructor() MaxStack {
	return MaxStack{}
}

func (s *MaxStack) Push(x int) {
	if 0 == len(s.maxIdx) || x >= s.PeekMax() {
		s.maxIdx = append(s.maxIdx, len(s.stack))
	}
	s.stack = append(s.stack, x)
}

func (s *MaxStack) Pop() int {
	lastIdx := len(s.stack) - 1
	ans := s.stack[lastIdx]
	s.stack = s.stack[:lastIdx]
	if s.maxIdx[len(s.maxIdx)-1] == lastIdx {
		s.maxIdx = s.maxIdx[:len(s.maxIdx)-1]
	}
	return ans
}

func (s *MaxStack) Top() int {
	lastIdx := len(s.stack) - 1
	ans := s.stack[lastIdx]
	return ans
}

func (s *MaxStack) PeekMax() int {
	return s.stack[s.maxIdx[len(s.maxIdx)-1]]
}

func (s *MaxStack) PopMax() int {
	maxIdx := len(s.maxIdx) - 1
	idx := s.maxIdx[maxIdx]
	ans := s.stack[idx]
	s.maxIdx = s.maxIdx[:maxIdx]
	s.stack = append(s.stack[:idx], s.stack[idx+1:]...)
	for i := idx; i < len(s.stack); i++ {
		x := s.stack[i]
		if 0 == len(s.maxIdx) || x >= s.PeekMax() {
			s.maxIdx = append(s.maxIdx, i)
		}
	}

	return ans
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */
