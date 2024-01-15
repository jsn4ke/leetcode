package leatcode

func repeatLimitedString(s string, repeatLimit int) string {
	var (
		pairs ['z' - 'a' + 1]int
		np    [][2]int
	)
	for _, v := range s {
		pairs[v-'a']++
	}
	for i, v := range pairs {
		if 0 == v {
			continue
		}
		np = append(np, [2]int{i, v})
	}
	var (
		bs []byte
	)
	if last := len(np) - 1; last >= 0 {
		a, cur := byte(np[last][0])+'a', np[last][1]

	}
	return string(bs)
}

// func repeatLimitedString(s string, repeatLimit int) string {
// 	bs := []byte(s)
// 	sort.Slice(bs, func(i, j int) bool {
// 		return bs[i] > bs[j]
// 	})
// 	var (
// 		startIdx = 0
// 		length = len(bs)
// 	)
// 	for i := 0; i < length; i++ {
// 		if bs[i] != bs[startIdx] {
// 			startIdx = i
// 			continue
// 		}
// 		if i-startIdx > repeatLimit {
// 			find := false
// 			for j := i + 1; j < len(bs); j++ {
// 				if bs[i] != bs[j] {
// 					bs[i], bs[j] = bs[j], bs[i]
// 					find = true
// 					break
// 				}
// 			}
// 			if !find {
// 				length = i
// 				break
// 			}
// 			startIdx = i
// 		}
// 	}
// 	return string(bs[:length])
// }
