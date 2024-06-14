package leetcode_155_MinStack

type MinStack struct {
	sk  []int
	min []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.sk = append([]int{val}, s.sk...)
	if 0 == len(s.min) || val < s.min[0] {
		s.min = append([]int{val}, s.min...)
	} else {
		s.min = append([]int{s.min[0]}, s.min...)
	}
}

func (s *MinStack) Pop() {
	s.sk = s.sk[1:]
	s.min = s.min[1:]
}

func (s *MinStack) Top() int {
	return s.sk[0]
}

func (s *MinStack) GetMin() int {
	return s.min[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
