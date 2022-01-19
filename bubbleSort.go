package main

import "fmt"

func main() {

	var forSort = []int{3, 7, 4, 1, 8, 2, 3, 9, 32, 12, 5, 56, 18}
	fmt.Printf("your slice %v before sorting\n", forSort)
	fmt.Printf("your slice %v after sorting", Sorting(forSort))
}
func Sorting(data []int) []int {
	var length int = len(data) // variable equal slice size for looping through all items in the slice

	for k := 1; k <= length; k++ { // loop for checking order of number as many times as items in the slice
		for i := length - 1; i > 0; i-- { // inner loop

			if data[i-1] <= data[i] { // if statemant which compare two number one next to other

			} else if data[i-1] > data[i] {
				var first int = data[i-1]
				var second int = data[i]
				data[i-1] = second // technical variable for value reversing
				data[i] = first    // technical variable for value reversing

			}
		}
	}
	return data
}
