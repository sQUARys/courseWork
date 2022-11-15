package sorts

import "math/rand"

func BubbleSort(n []int) {
	var sorted = false

	for !sorted {
		sorted = true
		i := 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i] //swap two elements
				sorted = false              // arr not sorted
			}
			i++ // add index
		}
	}

}

func InsertionSort(n []int) {
	var i = 1
	for i < len(n) {
		var j = i
		for j >= 1 && n[j] < n[j-1] { // shift the value until it is more
			n[j], n[j-1] = n[j-1], n[j] // swap two elements
			j--                         //reducing the index for the previous element
		}
		i++
	}
}

func SelectionSort(n []int) {
	i := 1

	for i < len(n)-1 {
		j := i + 1
		minIndex := i

		if j < len(n) { // find a min value
			if n[j] < n[minIndex] {
				minIndex = j
			}
			j++
		}

		if minIndex != i { // just start to swap values to the beginning of array
			var temp = n[i]
			n[i] = n[minIndex]
			n[minIndex] = temp
		}

		i++
	}

}

func Quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1 // set left and right side

	pivot := rand.Int() % len(a) // set random index into array

	a[pivot], a[right] = a[right], a[pivot] // swap right side and a random index

	for i, _ := range a { // sorting a pivot's side of arr
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left] // change sides of arr by place

	Quicksort(a[:left])   // recursion for elements staying before left index
	Quicksort(a[left+1:]) // recursion for elements staying after left index

	return a
}

func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2 // get middle part of array

	return MergeArrays(MergeSort(slice[:mid]), MergeSort(slice[mid:])) // we just merge two part of array: before and after middle index
}

func MergeArrays(left []int, right []int) []int {
	//Merge two slices in ascending order.
	result := make([]int, 0)              // create result slice
	for len(left) > 0 && len(right) > 0 { // for left and right not zero values
		if left[0] < right[0] {
			result = append(result, left[0]) // added two array a left value
			left = left[1:]                  // cut value which we already wrote
		} else {
			result = append(result, right[0]) // // added two array a left value
			right = right[1:]                 // cut valuСe which we already wrote
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}

func ShellSort(arr []int) []int {
	gap := len(arr) / 2

	for gap > 0 {
		for j := gap; j < len(arr); { // check arr from left to right
		loop: // just name of i-loop for beautiful stopping
			for i := j - gap; i >= 0; { // j keep help in maintain gap value
				//If value on right side is already greater than left side value
				// We don't do swap else we swap
				if arr[i+gap] > arr[i] {
					break loop
				} else {
					arr[i+gap], arr[i] = arr[i], arr[i+gap] // swap two value
				}
				i -= gap // To check left side also if the element present is greater than current element
			}
			j++
		}
		gap /= 2
	}
	return arr
}
