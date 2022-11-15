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
}

func New() *Service {
	return &Service{
		Numbers: []int{},
	}
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
