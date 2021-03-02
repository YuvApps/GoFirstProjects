package mediumLeetCode

import "math"

func MinSteps(n int) int {
	if n == 1 { return 0 }
	greatestPrime := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			isPrime := true
			for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
				if i % j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				greatestPrime = i
			}
		}
	}
	isMulti := true
	drill := n
	index := 0
	for index = 0; drill != 1 && isMulti;  {
		if drill % greatestPrime != 0 || greatestPrime == 1 {
			isMulti = false
		} else {
			drill /= greatestPrime
			index++
		}
	}

	retAnswer := greatestPrime + (n / greatestPrime)

	if isMulti {
		retAnswer = greatestPrime*index
	}

	if greatestPrime == 1 { retAnswer-- }
	return retAnswer
}
