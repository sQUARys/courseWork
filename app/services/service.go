package services

import (
	"courseWork/app/sorts"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Service struct {
	Numbers       []int
	SortedNumbers []int
	Sorts         Sorts
	*sync.RWMutex
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
}

func New(s Sorts) *Service {
	return &Service{
		Numbers: []int{},
		Sorts:   s,
	}
}

func (s *Service) SetArrayByUserChoice(w http.ResponseWriter, choice interface{}, choicesOfSorts []string) {
	switch choice.(type) {
	case string:
		s.FillFromFile(choice.(string))
	case int:
		s.FillByRand(choice.(int))
	}
	s.StartSorting(w, choicesOfSorts)
}

func (s *Service) StartSorting(w http.ResponseWriter, choicesOfSorts []string) {
	startedArray := s.Numbers

	var wg sync.WaitGroup

	typeOfSortChan := make(chan string)
	doneChan := make(chan interface{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, b := range choicesOfSorts {
			typeOfSortChan <- b
		}
		doneChan <- true
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.WorkerOfSort(startedArray, typeOfSortChan, doneChan)
	}()

	wg.Wait()
	close(typeOfSortChan)
	close(doneChan)
}

func (s *Service) WorkerOfSort(startedArray []int, sortsCh chan string, doneCh chan interface{}) {
loop:
	for {
		select {
		case v := <-sortsCh:
			switch v {
			case "Bubble":
				sortedArray := s.Sorts.BubbleSort(startedArray)
				if len(s.SortedNumbers) == 0 {
					s.SortedNumbers = sortedArray
				}
			case "Quick":
				sortedArray := s.Sorts.Quicksort(startedArray)
				if len(s.SortedNumbers) == 0 {
					s.SortedNumbers = sortedArray
				}
			case "Insertion":
				sortedArray := s.Sorts.InsertionSort(startedArray)
				if len(s.SortedNumbers) == 0 {
					s.SortedNumbers = sortedArray
				}
			case "Selection":
				sortedArray := s.Sorts.SelectionSort(startedArray)
				if len(s.SortedNumbers) == 0 {
					s.SortedNumbers = sortedArray
				}
			case "Merge":
				sortedArray := s.Sorts.MergeSort(startedArray)
				if len(s.SortedNumbers) == 0 {
					s.SortedNumbers = sortedArray
				}
			case "Shell":
				sortedArray := s.Sorts.ShellSort(startedArray)
				if len(s.SortedNumbers) == 0 {
					s.SortedNumbers = sortedArray
				}
			}
			fmt.Println(fmt.Sprintf("%s sorts done succesful", v))
		case <-doneCh:
			break loop
		}
	}

}

func (s *Service) FillByRand(n int) {
	s.Numbers = []int{}
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		s.Numbers = append(s.Numbers, rand.Intn(n))
	}
}

func (s *Service) FillFromFile(path string) error {
	s.Numbers = []int{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
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

func (s *Service) GetSortsResultJSON() string {
	data, _ := json.Marshal(s.Sorts)
	return string(data)
}

func (s *Service) GetStartedArray() []int {
	return s.Numbers
}

func (s *Service) GetAvailableSorts() []string {
	return s.Sorts.GetAvailableSorts()
}

func (s *Service) GetSortedArray() []int {
	return s.SortedNumbers
}

func (s *Service) CleanService() {
	s.Sorts = sorts.New()
}
