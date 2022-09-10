package main

import "fmt"

func main() {
	var input1, input2, x int
	input1 = 2
	input2 = 4
	x = 5
	arr := []int{input1, input2}
	selisih := input2 - input1

	current := input2
	for i := 0; i < (x - 2); i++ {
		current += selisih
		arr = append(arr, current)
	}
	fmt.Println(arr)
}
