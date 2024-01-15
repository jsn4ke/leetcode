package leatcode

func isOneEditDistance(s string, t string) bool {
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
	} else if len(s) < len(t) {
		var (
			si int
			ti int
		)
		for si := 0; si < len(s); si++ {
			if s[si] == t[ti] {
				ti++
				continue
			}
		}
	}
}
