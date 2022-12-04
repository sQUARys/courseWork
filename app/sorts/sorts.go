package sorts

import (
	"github.com/dorin131/go-data-structures/minheap"
	"math"
	"time"
)

type Sorts struct {
	Sorts []CertainSort
}

type CertainSort struct {
	Time       float64 `json:"time"`
	TypeOfSort string  `json:"type"`
}

var AvailableSort = []string{
	"Bubble",
	"Quick",
	"Selection",
	"Insertion",
	"Merge",
	"Shell",
	"Intro",
	"Tim",
}

func New() *Sorts {
	return &Sorts{
		Sorts: []CertainSort{},
	}
}

func (s *Sorts) AddSort(startTime time.Time, typeOfSort string) {
	result := CertainSort{Time: time.Since(startTime).Seconds(), TypeOfSort: typeOfSort}
	s.Sorts = append(s.Sorts, result)
}

func (s *Sorts) GetAvailableSorts() []string { // getter of all avialable sorts
	return AvailableSort
}

func (s *Sorts) CopyArr(n []int) []int { // special func to delete dependencies between func and arr
	tmp := make([]int, len(n))
	copy(tmp, n)
	return tmp
}

func (s *Sorts) BubbleSort(startedArray []int) []int {
	arrayForSort := s.CopyArr(startedArray) // array for sort

	startTime := time.Now() // start time ticker

	var sorted = false // checker of sorted arr

	for !sorted {
		sorted = true
		for i := 0; i < len(arrayForSort)-1; i++ {
			if arrayForSort[i] > arrayForSort[i+1] { // sorting
				arrayForSort[i], arrayForSort[i+1] = arrayForSort[i+1], arrayForSort[i] //swap two elements
				sorted = false                                                          // arr not sorted
			}
		}
	}

	s.AddSort(startTime, "Bubble") // just add sort info for html

	return arrayForSort
}

func (s *Sorts) InsertionSort(startedArray []int) []int {
	arrayForSort := s.CopyArr(startedArray) // delete dependencies

	startTime := time.Now() // start ticker

	for i := 1; i < len(arrayForSort); i++ {
		for j := i; j > 0; j-- {
			if arrayForSort[j-1] > arrayForSort[j] { // if more
				arrayForSort[j-1], arrayForSort[j] = arrayForSort[j], arrayForSort[j-1] // swap
			}
		}
	}

	s.AddSort(startTime, "Insertion") // just add sort info for html

	return arrayForSort
}

func (s *Sorts) SelectionSort(startedArray []int) []int {
	arrayForSort := s.CopyArr(startedArray)

	startTime := time.Now()

	for i := 0; i < len(arrayForSort); i++ { // run by all array by i
		var minIdx = i                           // minimal index
		for j := i; j < len(arrayForSort); j++ { // run by array by j
			if arrayForSort[j] < arrayForSort[minIdx] {
				minIdx = j
			}
		}
		arrayForSort[i], arrayForSort[minIdx] = arrayForSort[minIdx], arrayForSort[i] // swap
	}

	s.AddSort(startTime, "Selection")

	return arrayForSort
}

func (s *Sorts) Quicksort(startedArray []int) []int {
	startTime := time.Now() // start ticker
	copiedArray := s.CopyArr(startedArray)
	s.QuickSortRecursive(copiedArray, 0, len(copiedArray)-1)
	s.AddSort(startTime, "Quick") // add sort
	return copiedArray
}

func (s *Sorts) QuickSortRecursive(copiedArray []int, low int, high int) []int {
	if low < high {
		p := Partition(copiedArray, low, high)

		copiedArray = s.QuickSortRecursive(copiedArray, low, p-1)
		copiedArray = s.QuickSortRecursive(copiedArray, p+1, high)
	}
	return copiedArray
}

func (s *Sorts) MergeSort(startedArray []int) []int {
	startTime := time.Now()                           // start ticker
	sortedArray := s.MergeSortRecursive(startedArray) // start recursion part of mergesort
	s.AddSort(startTime, "Merge")
	return sortedArray
}

func (s *Sorts) MergeSortRecursive(startedArray []int) []int {
	slice := s.CopyArr(startedArray) // delete dependencies

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
	startTime := time.Now()                           // start ticker
	sortedArray := s.ShellSortRecursive(startedArray) // start recursion part of ShellSort
	s.AddSort(startTime, "Shell")
	return sortedArray
}

func (s *Sorts) ShellSortRecursive(startedArray []int) []int {
	arr := s.CopyArr(startedArray)

	gap := len(arr) / 2

	for gap > 0 {
		for j := gap; j < len(arr); j++ { // check arr from left to right
		loop: // just name of i-loop for beautiful stopping
			for i := j - gap; i >= 0; i -= gap { // j keep help in maintain gap value
				//If value on right side is already greater than left side value
				// We don't do swap else we swap
				if arr[i+gap] > arr[i] {
					break loop
				} else {
					arr[i+gap], arr[i] = arr[i], arr[i+gap] // swap two value
				}
			}
		}
		gap /= 2
	}
	return arr
}

func (s *Sorts) TimSort(startedArray []int) []int {
	startTime := time.Now() // start ticker

	n := len(startedArray)  // get len of arr
	minRun := CalcMinRun(n) // calculate minimal running
	copiedArray := s.CopyArr(startedArray)

	for start := 0; start < n; start += minRun {
		end := math.Min(float64(start+minRun-1), float64(n-1))
		//InsertionSort
		for i := start + 1; i < int(end)+1; i++ {
			for j := i; j > start && copiedArray[j] < copiedArray[j-1]; j-- {
				copiedArray[j], copiedArray[j-1] = copiedArray[j-1], copiedArray[j]
			}
		}
	}

	size := minRun
	for size < n {

		for left := 0; left < n; left += 2 * size {
			middle := math.Min(float64(n-1), float64(left+size-1))
			right := math.Min(float64(left+2*size-1), float64(n-1))

			if middle < right {
				MergeParts(copiedArray, left, int(middle), int(right))
			}
		}
		size *= 2
	}

	s.AddSort(startTime, "Tim")

	return copiedArray
}

func MergeParts(copiedArray []int, l int, m int, r int) {
	//merge(arr, left, mid, right)

	len1, len2 := m-l+1, r-m
	left, right := []int{}, []int{}

	for i := 0; i < len1; i++ {
		left = append(left, copiedArray[l+i])
	}
	for i := 0; i < len2; i++ {
		right = append(right, copiedArray[m+1+i])
	}
	i, j, k := 0, 0, l

	for i < len1 && j < len2 {
		if left[i] <= right[j] {
			copiedArray[k] = left[i]
			i++
		} else {
			copiedArray[k] = right[j]
			j++
		}
		k++
	}

	for i < len1 {
		copiedArray[k] = left[i]
		k++
		i++
	}
	for j < len2 {
		copiedArray[k] = right[j]
		k++
		j++
	}
}

func CalcMinRun(n int) int {
	r := 0
	for n >= 32 {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func (s *Sorts) IntroSort(startedArray []int) []int {
	startTime := time.Now()

	begin := 0
	end := len(startedArray) - 1
	arrayForSort := s.CopyArr(startedArray)

	depthLimit := 2 * math.Floor(math.Log2(float64(end)-float64(begin)))

	result := CertainSort{Time: time.Since(startTime).Seconds(), TypeOfSort: "Intro"}
	s.Sorts = append(s.Sorts, result)

	return s.IntroSortUtil(arrayForSort, begin, end, int(depthLimit))
}

func (s *Sorts) IntroSortUtil(copiedArray []int, begin int, end int, depthLimit int) []int {
	size := end - begin

	if size < 16 {
		//InsertionSort
		for i := 1; i < len(copiedArray); i++ {
			key := copiedArray[i]
			j := i - 1
			for j >= 0 && copiedArray[j] > key {
				copiedArray[j+1] = copiedArray[j]
				j--
			}
			copiedArray[j+1] = key
		}
		return copiedArray
	}

	if depthLimit == 0 {
		return HeapSort(copiedArray) // Heap sort
	}

	pivot := MedianOfThree(copiedArray, begin, begin+size/2, end)               // get median of three elem
	copiedArray[pivot], copiedArray[end] = copiedArray[end], copiedArray[pivot] // swap end and pivot

	partitionIndex := Partition(copiedArray, begin, end) // index
	s.IntroSortUtil(copiedArray, begin, partitionIndex-1, depthLimit-1)
	s.IntroSortUtil(copiedArray, partitionIndex+1, end, depthLimit-1)
	return copiedArray
}

func MedianOfThree(startedArray []int, first int, second int, third int) int { // getting middle of three
	firstElement := startedArray[first]
	secondElement := startedArray[second]
	thirdElement := startedArray[third]

	if firstElement <= secondElement && secondElement <= thirdElement {
		return second
	}
	if thirdElement <= secondElement && secondElement <= firstElement {
		return second
	}
	if secondElement <= firstElement && firstElement <= thirdElement {
		return first
	}
	if thirdElement <= firstElement && firstElement <= secondElement {
		return first
	}
	return third
}

func HeapSort(input []int) []int { // sort by heap
	result := []int{}

	mh := minheap.New(input)

	for range input {
		result = append(result, mh.ExtractMin())
	}

	return result
}

func Partition(copiedArray []int, low int, high int) int {
	pivot := copiedArray[high] // get high value

	i := low

	for j := low; j < high; j++ { // go for all elements in part
		if copiedArray[j] < pivot {
			copiedArray[i], copiedArray[j] = copiedArray[j], copiedArray[i] // swap elements
			i++
		}
	}

	copiedArray[i], copiedArray[high] = copiedArray[high], copiedArray[i] // swap elements
	return i
}
