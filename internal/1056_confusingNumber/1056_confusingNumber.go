package leetcode_1056_confusingNumber

func confusingNumber(n int) bool {
	hash := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}
	var ans int
	for tmp := n; tmp != 0; tmp /= 10 {
		part := tmp % 10
		if _, ok := hash[part]; !ok {
			return false
		}
		part = hash[part]
		ans = ans*10 + part
	}
	return ans != n
}
