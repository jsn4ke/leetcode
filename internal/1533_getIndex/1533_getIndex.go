package leetcode_1533_getIndex

// This is the ArrayReader's API interface.
// You should not implement it, or speculate about its implementation
type ArrayReader struct {
}

// Compares the sum of arr[l..r] with the sum of arr[x..y]
// return 1 if sum(arr[l..r]) > sum(arr[x..y])
// return 0 if sum(arr[l..r]) == sum(arr[x..y])
// return -1 if sum(arr[l..r]) < sum(arr[x..y])
func (this *ArrayReader) compareSub(l, r, x, y int) int {
	return 0
}

// Returns the length of the array
func (this *ArrayReader) length() int {
	return 1
}

func getIndex(reader *ArrayReader) int {
	var dfs func(s, t int) int
	dfs = func(s, t int) int {
		if s == t {
			return s
		}
		if (t-s)&1 == 0 {
			mid := s + (t-1-s)>>1
			val := reader.compareSub(s, mid, mid+1, t-1)
			switch val {
			case 0:
				return t
			case 1:
				return dfs(s, mid)
			case -1:
				return dfs(mid+1, t-1)
			}
		} else {
			mid := s + (t-s)>>1
			val := reader.compareSub(s, mid, mid+1, t)
			switch val {
			case 1:
				return dfs(s, mid)
			case -1:
				return dfs(mid+1, t)
			}
		}
		return -1
	}
	return dfs(0, reader.length()-1)
}
