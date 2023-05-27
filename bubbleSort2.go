package main

import "fmt"

func main() {
	n := []int{2, 3, 4, 8, 7, 9, 11, 13}
	fmt.Println(n)
	bubbleSort2(n)
	return
}

func bubbleSort2(numbers []int) {
	N := len(numbers)
	for i := 0; i < N; i++ {
		sweep2(numbers, i)
		fmt.Println(numbers)
	}

}

func sweep2(numbers []int, prevPass int) {
	N := len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	for secondIndex < N-prevPass {
		if numbers[firstIndex] > numbers[secondIndex] {
			numbers[firstIndex], numbers[secondIndex] = numbers[secondIndex], numbers[firstIndex]
		}
		firstIndex++
		secondIndex++
	}
	fmt.Println(firstIndex, secondIndex)
}
