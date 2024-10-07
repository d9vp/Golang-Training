package main

import "fmt"

func main() {
	num := 10
	check_prime(num)
	sum := sum_fibo(num)
	fmt.Printf("\nSum of %d elements in Fibonacci series is %d", num, sum)

	slc := []int{2, 0, 5, 7, 9, 0, 14, 27, 1}
	count_oez(slc)
}

func check_prime(num int) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Printf("%d is not prime.\n", num)
			return
		}
	}
	fmt.Printf("%d is prime.\n", num)
}

func sum_fibo(num int) int {
	fibo := []int{0, 1}
	sum := 1
	for i := 0; i < num-2; i++ {
		fibo = append(fibo, fibo[len(fibo)-2]+fibo[len(fibo)-1])
		sum += fibo[len(fibo)-1]
	}
	// fmt.Println(fibo, sum)
	return sum
}

func count_oez(slc []int) {
	zeroes, evens, odds := 0, 0, 0

	for _, val := range slc {
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
