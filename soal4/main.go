package main

import "fmt"

func main() {
	numbers := []float64{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	fmt.Println(BubbleSortAsc(numbers))
	fmt.Println(BubbleSortDesc(numbers))
}

func BubbleSortAsc(numbers []float64) []float64 {
	for i := 0; i < (len(numbers) - 1); i++ {
		for j := 0; j < (len(numbers) - 1); j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
		// fmt.Println("=", i, "=", numbers)
	}

	return numbers
}

func BubbleSortDesc(numbers []float64) []float64 {
	for i := 0; i < (len(numbers) - 1); i++ {
		for j := 0; j < (len(numbers) - 1); j++ {
			if numbers[j] < numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
		// fmt.Println("=", i, "=", numbers)
	}

	return numbers
}
