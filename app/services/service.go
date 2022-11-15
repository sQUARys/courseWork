package services

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	Numbers []int
	Sorts   Sorts
}

type Sorts interface {
	BubbleSort([]int)
	InsertionSort([]int)
	SelectionSort([]int)
	Quicksort([]int) []int
	MergeSort([]int) []int
	ShellSort([]int) []int
	CopyArr([]int) []int
}

func New(s Sorts) *Service {
	return &Service{
		Numbers: []int{},
		Sorts:   s,
	}
}

func (s *Service) StartSorting() {
	startedArray := s.Numbers

	s.Sorts.BubbleSort(startedArray)

}

func (s *Service) FillByRand(n int) {
	s.Numbers = []int{}
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		s.Numbers = append(s.Numbers, rand.Intn(100))
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
