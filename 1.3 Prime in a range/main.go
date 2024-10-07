package main

import "fmt"

func main() {
	primes := checkPrimeInRange(4, 40)
	fmt.Println(primes)
}

func checkPrimeInRange(num1 int, num2 int) []int {
	allPrimes := []int{}
	flag := 0
	for i := num1; i <= num2; i++ {
		flag = 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 {
			allPrimes = append(allPrimes, i)
		}

	}
	return allPrimes
}
