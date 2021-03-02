package easyLeetCode

func MaximumWealth(accounts [][]int) int {
	maxWealth := 0
	for personIndex := 0; personIndex < len(accounts); personIndex++ {
		personWealth := 0

		for accountIndex := 0; accountIndex < len(accounts[personIndex]); accountIndex++ {
			personWealth += accounts[personIndex][accountIndex]
		}

		if personWealth > maxWealth { maxWealth = personWealth }
	}
	return maxWealth
}
