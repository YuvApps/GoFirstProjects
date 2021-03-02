package main

import (
	"bufio"
	"firstProject/dailyChallengeLeetCode"
	"firstProject/easyLeetCode"
	"firstProject/mediumLeetCode"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	for level := 1; level != 0; {
		fmt.Println("1 - Easy")
		fmt.Println("2 - Medium")
		fmt.Println("3 - Hard")
		fmt.Println("4 - Daily Challenges")
		fmt.Println("Enter Level (0 to Exit):")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		level = num
		if err != nil {
			fmt.Println(err)
		}
		switch {
			case level == 1:
				for questNo := 1; questNo != 0; {
					fmt.Println("1 - TwoSum")
					fmt.Println("2 - RunningSum")
					fmt.Println("3 - DefangIPAddr")
					fmt.Println("4 - KidsWithCandies")
					fmt.Println("5 - Shuffle")
					fmt.Println("6 - MaximumWealth")
					fmt.Println("7 - NumIdenticalPairs")
					fmt.Println("8 - NumJewelsInStones")
					fmt.Println("Enter Question Number (0 to Exit):")
					scanner := bufio.NewScanner(os.Stdin)
					scanner.Scan()
					text := scanner.Text()
					num, err := strconv.Atoi(text)
					questNo = num
					if err != nil {
						fmt.Println(err)
					}
					switch {
					case questNo == 1:
						var nums []int
						for ; num != -1; {
							fmt.Println("Enter Number (-1 to Stop):")
							scanner.Scan()
							text = scanner.Text()
							num, err = strconv.Atoi(text)
							if num != -1 {
								nums = append(nums, num)
							}
						}
						fmt.Println("Enter Target:")
						scanner.Scan()
						text = scanner.Text()
						num, err = strconv.Atoi(text)
						target := num
						fmt.Println(easyLeetCode.TwoSum(nums, target))

					case questNo == 2:
						var nums2 []int
						for ; num != -1; {
							fmt.Println("Enter Number (-1 to Stop):")
							scanner.Scan()
							text = scanner.Text()
							num, err = strconv.Atoi(text)
							if num != -1 {
								nums2 = append(nums2, num)
							}
						}
						fmt.Println(easyLeetCode.RunningSum(nums2))

					case questNo == 3:
						fmt.Println("Enter Address:")
						scanner.Scan()
						address := scanner.Text()
						fmt.Println(easyLeetCode.DefangIPAddr(address))

					case questNo == 4:
						var nums4 []int
						for ; num != -1; {
							fmt.Println("Enter Candies (-1 to stop):")
							scanner.Scan()
							text = scanner.Text()
							num, err = strconv.Atoi(text)
							if num != -1 {
								nums4 = append(nums4, num)
							}
						}
						fmt.Println("Enter Extra Candies:")
						scanner.Scan()
						text = scanner.Text()
						num, err = strconv.Atoi(text)
						extraCandies := num
						fmt.Println(easyLeetCode.KidsWithCandies(nums4, extraCandies))

					case questNo == 5:
						var nums5 []int
						for ; num != -1; {
							fmt.Println("Enter Numbers (-1 to stop):")
							scanner.Scan()
							text = scanner.Text()
							num, err = strconv.Atoi(text)
							if num != -1 {
								nums5 = append(nums5, num)
							}
						}
						fmt.Println("Enter Number Of Half The List:")
						scanner.Scan()
						text = scanner.Text()
						num, err = strconv.Atoi(text)
						fmt.Println(easyLeetCode.Shuffle(nums5, num))

					case questNo == 6:
						var accounts [][]int
						fmt.Println("Enter Number People:")
						scanner.Scan()
						text = scanner.Text()
						num, err = strconv.Atoi(text)
						peopleNum := num
						fmt.Println("Enter Number Accounts Per Person:")
						scanner.Scan()
						text = scanner.Text()
						num, err = strconv.Atoi(text)
						accountsNum := num
						for personIndex := 0; personIndex < peopleNum; personIndex++ {
							accounts = append(accounts, []int{})
							for accountIndex := 0; accountIndex < accountsNum; accountIndex++ {
								fmt.Printf("Enter Account No. %d of Person %d ", accountIndex+1, personIndex+1)
								scanner.Scan()
								text = scanner.Text()
								num, err = strconv.Atoi(text)
								accounts[personIndex] = append(accounts[personIndex], num)
							}
						}
						fmt.Println(easyLeetCode.MaximumWealth(accounts))

					case questNo == 7:
						var nums7 []int
						for ; num != -1; {
							fmt.Println("Enter Number (-1 to Stop):")
							scanner.Scan()
							text = scanner.Text()
							num, err = strconv.Atoi(text)
							if num != -1 {
								nums7 = append(nums7, num)
							}
						}
						fmt.Println(easyLeetCode.NumIdenticalPairs(nums7))

					case questNo == 8:
						fmt.Println("Enter Jewels:")
						scanner.Scan()
						jewels := scanner.Text()
						fmt.Println("Enter Stones:")
						scanner.Scan()
						stones := scanner.Text()
						fmt.Println(easyLeetCode.NumJewelsInStones(jewels, stones))

					default:
						fmt.Println("Don't Have Such A Question.. Please Try Again")
					}
				}
			case level == 2:
				for questNo := 1; questNo != 0; {
					fmt.Println("1 - MinOperations")
					fmt.Println("2 - MinSteps")
					fmt.Println("Enter Question Number (0 to Exit):")
					scanner := bufio.NewScanner(os.Stdin)
					scanner.Scan()
					text := scanner.Text()
					num, err := strconv.Atoi(text)
					questNo = num
					if err != nil {
						fmt.Println(err)
					}
					switch {
						case questNo == 1:
							fmt.Println("Enter Boxes:")
							scanner.Scan()
							boxes := scanner.Text()
							fmt.Println(mediumLeetCode.MinOperations(boxes))
						case questNo == 2:
							fmt.Println("Enter Number Accounts Per Person:")
							scanner.Scan()
							text = scanner.Text()
							num, err = strconv.Atoi(text)
							fmt.Println(mediumLeetCode.MinSteps(num))
						default:
							fmt.Println("Don't Have Such A Question.. Please Try Again")
					}
				}
			case level == 4:
				for questNo := 1; questNo != 0; {
					fmt.Println("1 - DistributeCandies")
					fmt.Println("2 - FindErrorNums")
					fmt.Println("Enter Question Number (0 to Exit):")
					scanner := bufio.NewScanner(os.Stdin)
					scanner.Scan()
					text := scanner.Text()
					num, err := strconv.Atoi(text)
					questNo = num
					if err != nil {
						fmt.Println(err)
					}
					switch {
						case questNo == 1:
							var nums []int
							for ; num != -1; {
								fmt.Println("Enter Number (-1 to Stop):")
								scanner.Scan()
								text = scanner.Text()
								num, err = strconv.Atoi(text)
								if num != -1 {
									nums = append(nums, num)
								}
							}
							fmt.Println(dailyChallengeLeetCode.DistributeCandies(nums))
						case questNo == 2:
							var nums2 []int
							for ; num != -1; {
								fmt.Println("Enter Number (-1 to Stop):")
								scanner.Scan()
								text = scanner.Text()
								num, err = strconv.Atoi(text)
								if num != -1 {
									nums2 = append(nums2, num)
								}
							}
							fmt.Println(dailyChallengeLeetCode.FindErrorNums(nums2))
						default:
							fmt.Println("Don't Have Such A Question.. Please Try Again")
					}
				}
		}
	}
}