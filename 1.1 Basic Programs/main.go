package main

import "fmt"

func main() {
	num := 10
	checkPrime(num)
	sum := sumFibonacci(num)
	fmt.Printf("\nSum of %d elements in Fibonacci series is %d", num, sum)

	mySlice := []int{2, 0, 5, 7, 9, 0, 14, 27, 1}
	countEvenOddZero(mySlice)
}

func checkPrime(num int) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Printf("%d is not prime.\n", num)
			return
		}
	}
	fmt.Printf("%d is prime.\n", num)
}

func sumFibonacci(num int) int {
	fibo := []int{0, 1}
	sum := 1
	for i := 0; i < num-2; i++ {
		fibo = append(fibo, fibo[len(fibo)-2]+fibo[len(fibo)-1])
		sum += fibo[len(fibo)-1]
	}
	// fmt.Println(fibo, sum)
	return sum
}

func countEvenOddZero(mySlice []int) {
	zeroes, evens, odds := 0, 0, 0

	for _, val := range mySlice {
		if val == 0 {
			zeroes += 1
		} else if val%2 == 1 {
			odds += 1
		} else {
			evens += 1
		}
	}

	fmt.Printf("\n\nNumber of even numbers: %d.\n", evens)
	fmt.Printf("Number of odd numbers: %d.\n", odds)
	fmt.Printf("Number of zeroes: %d.\n", zeroes)
}
