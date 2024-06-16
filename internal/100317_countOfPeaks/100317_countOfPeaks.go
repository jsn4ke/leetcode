package leetcode_100317_countOfPeaks

func countOfPeaks(nums []int, queries [][]int) []int {
	st := NewSegmentTree(nums)
	result := []int{}

	for _, query := range queries {
		if query[0] == 1 {
			l, r := query[1], query[2]
			if r > l { // To handle the condition to ignore the first and last element
				result = append(result, st.Query(l+1, r-1))
			} else {
				result = append(result, 0)
			}
		} else if query[0] == 2 {
			st.Update(query[1], query[2])
		}
	}
	return result
}

type SegmentTree struct {
	data  []int
	peaks []int
	tree  []int
	n     int
}

func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)
	st := &SegmentTree{
		data:  make([]int, n),
		peaks: make([]int, n),
		tree:  make([]int, 4*n),
		n:     n,
	}
	copy(st.data, nums)
	for i := 1; i < n-1; i++ {
		st.peaks[i] = st.isPeak(i)
	}
	st.build(0, 0, n-1)
	return st
}

func (st *SegmentTree) isPeak(i int) int {
	if i > 0 && i < st.n-1 && st.data[i] > st.data[i-1] && st.data[i] > st.data[i+1] {
		return 1
	}
	return 0
}

func (st *SegmentTree) build(index, l, r int) {
	if l == r {
		st.tree[index] = st.peaks[l]
	} else {
		mid := (l + r) / 2
		st.build(2*index+1, l, mid)
		st.build(2*index+2, mid+1, r)
		st.tree[index] = st.tree[2*index+1] + st.tree[2*index+2]
	}
}

func (st *SegmentTree) Update(idx, val int) {
	if idx < 0 || idx >= st.n {
		return
	}
	st.update(idx, val)
	st.data[idx] = val
	st.refresh(idx - 1)
	st.refresh(idx)
	st.refresh(idx + 1)
}

func (st *SegmentTree) update(idx, val int) {
	if idx >= 0 && idx < st.n {
		st.data[idx] = val
	}
}

func (st *SegmentTree) refresh(idx int) {
	if idx < 1 || idx >= st.n-1 {
		return
	}
	newValue := st.isPeak(idx)
	if newValue != st.peaks[idx] {
		st.peaks[idx] = newValue
		st.updateTree(0, 0, st.n-1, idx)
	}
}

func (st *SegmentTree) updateTree(index, l, r, idx int) {
	if l == r {
		st.tree[index] = st.peaks[idx]
	} else {
		mid := (l + r) / 2
		if idx <= mid {
			st.updateTree(2*index+1, l, mid, idx)
		} else {
			st.updateTree(2*index+2, mid+1, r, idx)
		}
		st.tree[index] = st.tree[2*index+1] + st.tree[2*index+2]
	}
}

func (st *SegmentTree) Query(l, r int) int {
	return st.query(0, 0, st.n-1, l, r)
}

func (st *SegmentTree) query(index, sl, sr, l, r int) int {
	if sl > r || sr < l {
		return 0
	}
	if sl >= l && sr <= r {
		return st.tree[index]
	}
	mid := (sl + sr) / 2
	left := st.query(2*index+1, sl, mid, l, r)
	right := st.query(2*index+2, mid+1, sr, l, r)
	return left + right
}
