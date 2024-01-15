package leatcode

func isOneEditDistance(s string, t string) bool {
	if s == t {
		return false
	}
	if len(s) == len(t) {
		once := true
		for i := range s {
			if s[i] != t[i] {
				if !once {
					return false
				}
				once = false
			}
		}
	} else if len(s)+1 == len(t) {
		var (
			si   int
			ti   int
			once = true
		)
		for si = 0; si < len(s); si++ {
			if s[si] != t[ti] {
				if !once {
					return false
				}
				si--
				once = false
			}
			ti++
		}
	} else if len(t)+1 == len(s) {
		var (
			si   int
			ti   int
			once = true
		)
		for ti = 0; ti < len(t); ti++ {
			if s[si] == t[ti] {
				si++
				continue
			} else {
				if !once {
					return false
				}
				once = false
				si += 2
			}
		}
	}
	return true
}
