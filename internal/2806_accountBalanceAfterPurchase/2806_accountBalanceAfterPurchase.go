package leetcode_2806_accountBalanceAfterPurchase

func accountBalanceAfterPurchase(purchaseAmount int) int {
	if purchaseAmount%10 >= 5 {
		return 100 - (purchaseAmount/10*10 + 10)
	} else {
		return 100 - purchaseAmount/10*10
	}
}
