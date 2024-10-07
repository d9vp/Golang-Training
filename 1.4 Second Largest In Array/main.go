package main

import "fmt"

func main() {
	myArray := []int{45, 12, 78, 92, 34, -37, 89} //technically, used a slice instead of "array"
	secondLargest := secondLargestInArray(myArray)
	fmt.Printf("The second largest element in given array is %d.", secondLargest)
}

func secondLargestInArray(myArray []int) int {
	if len(myArray) >= 2 {
		largest, secondLargest := myArray[1], myArray[0]
		if myArray[0] > myArray[1] {
			largest, secondLargest = myArray[0], myArray[1]
		}

		for ind := range len(myArray) - 2 {
			if myArray[ind+2] > largest {
				secondLargest = largest
				largest = myArray[ind+2]
			} else if myArray[ind+2] > secondLargest {
				secondLargest = myArray[ind+2]
			}
		}
		return secondLargest
	} else {
		fmt.Println("You need atleast 2 elements to find the second largest in an array.")
		return -9999
	}

}
