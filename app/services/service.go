package services

import (
	"courseWork/app/sorts"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Service struct {
	Sorts Sorts
	*sync.RWMutex
	Numbers       []int
	SortedNumbers []int
}

type Sorts interface {
	BubbleSort([]int) []int
	InsertionSort([]int) []int
	SelectionSort([]int) []int
	Quicksort([]int) []int
	MergeSort([]int) []int
	ShellSort([]int) []int
	CopyArr([]int) []int
	GetAvailableSorts() []string
	TimSort(startedArray []int) []int
	IntroSort(startedArray []int) []int
}

//error not going to startsorting

func New(s Sorts) *Service {
	return &Service{
		Numbers:       []int{},
		SortedNumbers: []int{},
		Sorts:         s,
	}
}

func (s *Service) SetArrayByUserChoice(choice interface{}, choicesOfSorts []string) error {
	var err error
	switch choice.(type) { //depending on the type we choose how to add elem into array
	case string:
		err = s.FillFromFile(choice.(string)) // get data from file
	case int:
		s.FillByRand(choice.(int)) // set random value into array
	}
	if err != nil {
		return err
	}

	err = s.StartSorting(choicesOfSorts)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) StartSorting(choicesOfSorts []string) error {
	startedArray := s.Numbers // get started array

	var wg sync.WaitGroup // add waitgroup for goroutines

	typeOfSortChan := make(chan string) // channel of types of sorts
	doneChan := make(chan interface{})  // channel for closing all
	errorChan := make(chan error, 1)    // channel for error

	wg.Add(1) // add count of new goroutines
	go func() {
		defer wg.Done()
		for _, b := range choicesOfSorts { // go for all sorts which user chose
			typeOfSortChan <- b // write into chan
		}
		doneChan <- true // when we end to write into type chan we end our for loop
	}()

	var err error

	wg.Add(1) // add count of new goroutines
	go func() {
		defer wg.Done()
		err = s.WorkerOfSort(&wg, startedArray, typeOfSortChan, doneChan, errorChan) // start worker which will start all sorts
	}()

	wg.Wait() // wait closing of all goroutines
	if err != nil {
		return err
	}

	close(typeOfSortChan) //close all channels
	close(doneChan)
	close(errorChan)
	return nil
}

func (s *Service) WorkerOfSort(wg *sync.WaitGroup, startedArray []int, sortsCh chan string, doneCh chan interface{}, errCh chan error) error {
loop:
	for { // start infinite for loop
		select {
		case errFromCh := <-errCh:
			return errFromCh
		case v := <-sortsCh: // read from chan
			wg.Add(1)
			go func() {
				defer wg.Done()
				var sortedArray []int
				switch v { // start sorts
				case "Bubble":
					sortedArray = s.Sorts.BubbleSort(startedArray)
				case "Quick":
					sortedArray = s.Sorts.Quicksort(startedArray)
				case "Insertion":
					sortedArray = s.Sorts.InsertionSort(startedArray)
				case "Selection":
					sortedArray = s.Sorts.SelectionSort(startedArray)
				case "Merge":
					sortedArray = s.Sorts.MergeSort(startedArray)
				case "Shell":
					sortedArray = s.Sorts.ShellSort(startedArray)
				case "Intro":
					sortedArray = s.Sorts.IntroSort(startedArray)
				case "Tim":
					sortedArray = s.Sorts.TimSort(startedArray)
				}
				if len(s.SortedNumbers) == 0 && sort.IntsAreSorted(sortedArray) {
					s.SortedNumbers = sortedArray
				}
				s.CheckError(sortedArray, v, errCh)
				fmt.Println(fmt.Sprintf("%s sorts end", v))
			}()
		case <-doneCh:
			break loop
		}
	}
	return nil
}

func (s *Service) FillByRand(n int) {
	s.Numbers = []int{}
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		s.Numbers = append(s.Numbers, rand.Intn(10*n))
	}
}

func (s *Service) FillFromFile(path string) error {
	s.Numbers = []int{}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileText, err := ioutil.ReadAll(file)

	numbersString := strings.Fields(string(fileText))

	for _, number := range numbersString {
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			return err
		}
		s.Numbers = append(s.Numbers, numberInt)
	}

	return nil
}

func (s *Service) GetSortsResultJSON() (string, error) {
	data, err := json.Marshal(s.Sorts)
	return string(data), err
}

func (s *Service) GetStartedArray() []int {
	return s.Numbers
}

func (s *Service) GetSortedArray() []int {
	return s.SortedNumbers
}

func (s *Service) GetAvailableSorts() []string {
	return s.Sorts.GetAvailableSorts()
}

func (s *Service) CleanService() {
	s.Sorts = sorts.New()
	s.Numbers = []int{}
	s.SortedNumbers = []int{}
}

func (s *Service) CheckError(sortedArray []int, typeOfSort string, errCh chan error) {
	fmt.Println(reflect.DeepEqual(s.SortedNumbers, sortedArray), typeOfSort)
	if s.SortedNumbers != nil && reflect.DeepEqual(s.SortedNumbers, sortedArray) == false {
		errCh <- errors.New(fmt.Sprintf("%s Sorted array not equal with memory array", typeOfSort))
	}
}
