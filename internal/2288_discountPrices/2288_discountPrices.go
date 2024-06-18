package leetcode_2288_discountPrices

import "fmt"

func discountPrices(sentence string, discount int) string {
	var (
		ans string
	)
	for i := 0; i < len(sentence); i++ {
		for ; i < len(sentence); i++ {
			if sentence[i] != ' ' {
				break
			}
			ans += " "
		}
		var (
			j = i
		)
		for ; j < len(sentence); j++ {
			if sentence[j] == ' ' {
				break
			}
		}
		if i == j {
			continue
		}
		word := sentence[i:j]
		i = j - 1
		if len(word) > 1 && word[0] == '$' {
			var val float64
			for _, v := range word[1:] {
				if v >= '0' && v <= '9' {
					val = val*10 + float64(v-'0')
				} else {
					val = -1
					break
				}
			}
			if val != -1 {
				val = val * (1 - float64(discount)/100)
				word = fmt.Sprintf("$%.2f", val)
			}
		}
		ans += word
	}
	return ans
}
