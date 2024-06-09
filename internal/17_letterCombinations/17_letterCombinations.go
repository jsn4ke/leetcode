package leetcode_17_letterCombinations

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var (
		res   []string
		table = map[byte]string{
			'2': "abc",
			'3': "def",
			'4': "ghi",
			'5': "jkl",
			'6': "mno",
			'7': "pqrs",
			'8': "tuv",
			'9': "wxyz",
		}
	)
	var dfs func(int, string)
	dfs = func(idx int, path string) {
		if idx == len(digits) {
			res = append(res, path)
			return
		}
		for _, c := range table[digits[idx]] {
			dfs(idx+1, path+string(c))
		}
	}
	dfs(0, "")
	return res
}
