package sorts

import (
	"math/rand"
	"time"
)

type Sorts struct {
	Sorts []CertainSort
}

type CertainSort struct {
	Time       float64 `json:"time"`
	TypeOfSort string  `json:"type"`
}

var AvailableSort = []string{"Bubble", "Quick", "Selection", "Insertion", "Merge", "Shell"}

func New() *Sorts {
	return &Sorts{
		Sorts: []CertainSort{},
	}
}

func (s *Sorts) GetAvailableSorts() []string {
	return AvailableSort
}

func (s *Sorts) CopyArr(n []int) []int { // special func to delete dependencies between func and arr
	tmp := make([]int, len(n))
	copy(tmp, n)
	return tmp
}

func (s *Sorts) BubbleSort(startedArray []int) []int {
	arrayForSort := s.CopyArr(startedArray) // array for sort

	startTime := time.Now()

	var sorted = false

	for !sorted {
		sorted = true
		i := 0
		for i < len(arrayForSort)-1 {
			if arrayForSort[i] > arrayForSort[i+1] {
				arrayForSort[i], arrayForSort[i+1] = arrayForSort[i+1], arrayForSort[i] //swap two elements
				sorted = false                                                          // arr not sorted
			}
			i++ // add index
		}
	}

	result := CertainSort{Time: time.Since(startTime).Seconds(), TypeOfSort: "Bubble"}
	s.Sorts = append(s.Sorts, result)
	return arrayForSort
}

func (s *Sorts) InsertionSort(startedArray []int) []int {
	arrayForSort := s.CopyArr(startedArray)

	startTime := time.Now()

	var i = 1
	for i < len(arrayForSort) {
		var j = i
		for j >= 1 && arrayForSort[j] < arrayForSort[j-1] { // shift the value until it is more
			arrayForSort[j], arrayForSort[j-1] = arrayForSort[j-1], arrayForSort[j] // swap two elements
			j--                                                                     //reducing the index for the previous element
		}
		i++
	}

	result := CertainSort{Time: time.Since(startTime).Seconds(), TypeOfSort: "Insertion"}
	s.Sorts = append(s.Sorts, result)
	return arrayForSort
}

func (s *Sorts) SelectionSort(startedArray []int) []int {
	arrayForSort := s.CopyArr(startedArray)

	startTime := time.Now()

	i := 1

	for i < len(arrayForSort)-1 {
		j := i + 1
		minIndex := i

		if j < len(arrayForSort) { // find a min value
			if arrayForSort[j] < arrayForSort[minIndex] {
				minIndex = j
			}
			j++
		}

		if minIndex != i { // just start to swap values to the beginning of array
			var temp = arrayForSort[i]
			arrayForSort[i] = arrayForSort[minIndex]
			arrayForSort[minIndex] = temp
		}

		i++
	}

	result := CertainSort{Time: time.Since(startTime).Seconds(), TypeOfSort: "Selection"}
	s.Sorts = append(s.Sorts, result)
	return arrayForSort
}

func (s *Sorts) Quicksort(startedArray []int) []int {
	startTime := time.Now()
	sortedArray := s.QuickSortRecursive(startedArray)
	result := CertainSort{Time: time.Since(startTime).Seconds(), TypeOfSort: "Quick"}
	s.Sorts = append(s.Sorts, result)
	return sortedArray
}

func (s *Sorts) QuickSortRecursive(startedArray []int) []int {
	a := s.CopyArr(startedArray)

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

	s.QuickSortRecursive(a[:left])   // recursion for elements staying before left index
	s.QuickSortRecursive(a[left+1:]) // recursion for elements staying after left index

	return a
}

func (s *Sorts) MergeSort(startedArray []int) []int {
	startTime := time.Now()
	sortedArray := s.MergeSortRecursive(startedArray)
	result := CertainSort{Time: time.Since(startTime).Seconds(), TypeOfSort: "Merge"}
	s.Sorts = append(s.Sorts, result)
	return sortedArray
}

func (s *Sorts) MergeSortRecursive(startedArray []int) []int {
	slice := s.CopyArr(startedArray)

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2 // get middle part of array

	return MergeArrays(s.MergeSortRecursive(slice[:mid]), s.MergeSortRecursive(slice[mid:])) // we just merge two part of array: before and after middle index
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

func (s *Sorts) ShellSort(startedArray []int) []int {
	startTime := time.Now()
	sortedArray := s.ShellSortRecursive(startedArray)
	result := CertainSort{Time: time.Since(startTime).Seconds(), TypeOfSort: "Shell"}
	s.Sorts = append(s.Sorts, result)
	return sortedArray
}

func (s *Sorts) ShellSortRecursive(startedArray []int) []int {
	arr := s.CopyArr(startedArray)

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
