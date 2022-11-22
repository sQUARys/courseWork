package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Controller struct {
	Service service
	sync.RWMutex
}

type Times struct {
	TimesOfSorts []TimeOfSort
}

type TimeOfSort struct {
	SortType     int     `json:"type"`
	TimeDuration float64 `json:"time"`
}

type service interface {
	StartSorting()
	FillByRand(n int)
	FillFromFile(path string) error
	SetArrayByUserChoice(choice interface{})
	GetSortsResultJSON() string
}

func New(service service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (ctr *Controller) SendUserChoice(w http.ResponseWriter, r *http.Request) {
	//здесь будет еще валидация

	if filename := r.PostFormValue("filenameArr"); filename != "" {
		ctr.Service.SetArrayByUserChoice(filename)
	} else if sizeRandArr := r.PostFormValue("sizeRandArr"); sizeRandArr != "" {
		size, err := strconv.Atoi(sizeRandArr)
		if err != nil {
			log.Fatal(err)
			return
		}
		ctr.Service.SetArrayByUserChoice(size)
	}

	dataJSON := ctr.Service.GetSortsResultJSON()
	var dataStruct map[string][]TimeOfSort
	json.Unmarshal([]byte(dataJSON), &dataStruct)

	result := Times{dataStruct["Sorts"]}
	fmt.Println(result)
	err := WriteHTML(w, "app/templates/choiceMenu.html", result)
	if err != nil {
		log.Println(fmt.Sprintf("Error in  writing html. %w", err))
	}
}

func (ctr *Controller) GetSorts(w http.ResponseWriter, r *http.Request) {
	err := WriteHTML(w, "app/templates/choiceMenu.html", nil)
	if err != nil {
		log.Println(fmt.Sprintf("Error in  writing html. %w", err))
	}
}

func WriteHTML(w http.ResponseWriter, filename string, data interface{}) error {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return err
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return err
	}
	return nil
}
