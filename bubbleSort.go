package bubbleSort

func Sorting(data []int) []int {
	//standard bubble sorting algorithm. Sort slice of integers from small to bigger
	var length int = len(data) // variable equal slice size for looping through all items in the slice

	for k := 1; k <= length; k++ { // loop for checking order of number as many times as items in the slice
		for i := length - 1; i > 0; i-- { // inner loop

			if data[i-1] <= data[i] { // if statemant which compare two number one next to other

			} else if data[i-1] > data[i] {
				var first int = data[i-1]
				var second int = data[i]
				data[i-1], data[i] = second, first // technical variables for value reversing

			}
		}
	}
	return data
}

func ReverseSort(data []int) []int {
	//standard bubble sorting algorithm. Sort slice of integers from bigger to small
	var length int = len(data) // variable equal slice size for looping through all items in the slice

	for k := 1; k <= length; k++ { // loop for checking order of number as many times as items in the slice
		for i := 0; i <= length-2; i++ { // inner loop

			if data[i] >= data[i+1] { // if statemant which compare two number one next to other

			} else if data[i] < data[i+1] {
				var first int = data[i+1]
				var second int = data[i]
				data[i+1], data[i] = second, first // technical variables for value reversing

			}
		}
	}
	return data

}

func BubbleRecurtion(data []int) []int {
	// int slice sorting with recurtion
	counter := 0
	if len(data) == 1 || len(data) == 0 {
		return data
	} else {
		for i := 0; i < len(data)-1; i++ {
			for j := 0; j < len(data)-1; j++ {
				if data[j] > data[j+1] {
					data[j], data[j+1] = data[j+1], data[j]
				}
			}
		}
	}
	counter += 1
	BubbleRecurtion(data[counter:])
	return data
}
