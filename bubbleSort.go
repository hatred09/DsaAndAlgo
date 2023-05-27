package main

import (
	"fmt"
)

/*
Optimized bubble sort implementation
1. Don’t compare numbers that are already in their final position.
2. Stop sorting if our list is already sorted.
*/
func main() {
	n := []int{2, 3, 4, 8, 7, 9, 11, 13}
	fmt.Println(n)
	bubbleSort(n)
	return
}

func bubbleSort(numbers []int) {
	var N = len(numbers)
	for i := 0; i < N; i++ {
		flag := sweep(numbers, i)
		fmt.Println(numbers)
		if !flag {
			return
		}
	}

}

func sweep(numbers []int, prevPass int) bool {
	N := len(numbers)
	var flag bool
	var firstIndex = 0
	var secondIndex = 1
	for secondIndex < N-prevPass {
		if numbers[firstIndex] > numbers[secondIndex] {
			numbers[firstIndex], numbers[secondIndex] = numbers[secondIndex], numbers[firstIndex]
			flag = true
		}
		firstIndex++
		secondIndex++
	}
	fmt.Println(firstIndex, secondIndex)
	return flag
}
