package main

import "fmt"

func main() {
	n := []int{2, 3, 4, 8, 7, 9, 11, 13}
	fmt.Println(n)
	bubbleSort1(n)
	return
}

func bubbleSort1(numbers []int) {
	var N int = len(numbers)
	for i := 0; i < N; i++ {
		flag := sweep1(numbers)
		fmt.Println(numbers)
		if !flag {
			return
		}
	}

}

func sweep1(numbers []int) bool {
	var flag bool
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	for secondIndex < N {
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
