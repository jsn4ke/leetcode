package leatcode

func canReorderDoubled(arr []int) bool {
	mp := make(map[int]int, len(arr))
	for _, v := range arr {
		mp[v]++
	}
	for k, v := range mp {
		double := 2 * k
		// 奇数
		if 1 == k&1 {
			num := mp[double]
			if num < v {
				return false
			} else if num == v {
				delete(mp, k)
				delete(mp, 2*k)
			} else {
				mp[2*k] -= v
			}
		} else {

		}
	}
}
