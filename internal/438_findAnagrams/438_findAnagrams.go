package leetcode_438_findAnagrams

func findAnagrams(s string, p string) []int {
	var (
		res   []int
		aim   ['z' - 'a']int
		cache ['z' - 'a']int
	)
	if len(s) < len(p) {
		return nil
	}
	for i := range p {
		aim[p[i]-'a']++
		cache[s[i]-'a']++
	}
	if aim == cache {
		res = append(res, 0)
	}
	for i := range s[len(p):] {
		cache[s[i]-'a']--
		cache[s[i+len(p)]-'a']++
		if aim == cache {
			res = append(res, i+1)
		}
	}
	return res
}

func findAnagrams2(s string, p string) []int {
	var (
		res         []int
		cache           = make(map[byte]int)
		aim             = make(map[byte]int)
		left, right int = 0, 0
		size            = len(s)
	)
	for i := range p {
		aim[p[i]]++
	}
	for right < size {
		for right-left < len(p) && right < size {
			cache[s[right]]++
			right++
		}
		if len(cache) == len(aim) {
			ok := true
			for k, v := range aim {
				if cache[k] != v {
					ok = false
					break
				}
			}
			if ok {
				res = append(res, left)
			}
		}
		cache[s[left]]--
		if cache[s[left]] == 0 {
			delete(cache, s[left])
		}
		left++
	}
	return res
}
