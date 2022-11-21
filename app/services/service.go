package services

import (
	"fmt"
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
	Quicksort([]int)
	MergeSort([]int)
	ShellSort([]int)
	CopyArr([]int) []int
}

func New(s Sorts) *Service {
	return &Service{
		Numbers: []int{},
		Sorts:   s,
	}
}

func (s *Service) SetArrayByUserChoice(choice interface{}) {
	switch choice.(type) {
	case string:
		s.FillFromFile(choice.(string))
	case int:
		s.FillByRand(choice.(int))
	}
	s.StartSorting()
}

func (s *Service) StartSorting() {

	startedArray := s.Numbers
	fmt.Println("Started array: ", startedArray)

	s.Sorts.BubbleSort(startedArray)
	s.Sorts.SelectionSort(startedArray)
	s.Sorts.InsertionSort(startedArray)
	s.Sorts.Quicksort(startedArray)
	s.Sorts.MergeSort(startedArray)
	s.Sorts.ShellSort(startedArray)
	fmt.Println(s.Sorts)
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
