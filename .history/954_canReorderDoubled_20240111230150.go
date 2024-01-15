package leatcode

func canReorderDoubled(arr []int) bool {
	mp := make(map[int]int, len(arr))
	for _, v := range arr {
		mp[v]++
	}
loop:
	for k, v := range mp {
		double := 2 * k
		// 奇数
		if 1 == k&1 {
			// 肯定是小数
			num := mp[double]
			if num < v {
				return false
			} else if num == v {
				delete(mp, k)
				delete(mp, double)
			} else {
				mp[double] -= v
			}
		} else {
			// 大数
			num := mp[double]

			if num < v {
				// 大数不够
				little := v - num
				mp[k/2] -= little
				// 找小数
				if mp[k/2] < 0 {
					return false
				} else if mp[k/2] == 0 {
					delete(mp, k/2)
				}
				mp[k] -= little
			}
			// 小数
			num = mp[k/2]
		}
	}
	if 0 == len(mp) {
		return true
	} else {
		goto loop
	}
}
