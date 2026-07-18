func distMoney(money int, children int) int {
    if money < children {
		return -1
	}

	money -= children
	ans := money / 7
	if ans > children {
		ans = children
	}

	money -= ans * 7
	children -= ans

	if (children == 0 && money > 0) || (children == 1 && money == 3) {
		ans--
	}

	return ans
}