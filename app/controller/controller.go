package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Controller struct {
	Service service
	sync.RWMutex
}

type Times struct {
	AvailableSorts []string
	StartedArray   []int
	SortedArray    []int
	TimesOfSorts   []TimeOfSort
}

type TimeOfSort struct {
	SortType     string  `json:"type"`
	TimeDuration float64 `json:"time"`
}

type service interface {
	StartSorting(http.ResponseWriter, []string)
	FillByRand(n int)
	FillFromFile(path string) error
	SetArrayByUserChoice(http.ResponseWriter, interface{}, []string)
	GetSortsResultJSON() string
	GetStartedArray() []int
	CleanService()
	GetAvailableSorts() []string
	GetSortedArray() []int
}

func New(service service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (ctr *Controller) SendUserChoice(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//здесь будет еще валидация

	var input interface{}

	if filename := r.PostFormValue("filenameArr"); filename != "" {
		input = filename
	} else if sizeRandArr := r.PostFormValue("sizeRandArr"); sizeRandArr != "" {
		size, err := strconv.Atoi(sizeRandArr)
		if err != nil {
			log.Fatal(err)
			return
		}
		input = size
	}

	choicesOfSorts := r.PostForm["checkbox"]
	ctr.Service.SetArrayByUserChoice(w, input, choicesOfSorts)

	dataJSON := ctr.Service.GetSortsResultJSON()
	var dataStruct map[string][]TimeOfSort
	err := json.Unmarshal([]byte(dataJSON), &dataStruct)
	if err != nil {
		log.Println(err)
		return
	}

	result := Times{
		StartedArray:   ctr.Service.GetStartedArray(),
		AvailableSorts: ctr.Service.GetAvailableSorts(),
		SortedArray:    ctr.Service.GetSortedArray(),
		TimesOfSorts:   dataStruct["Sorts"],
	}

	err = WriteHTML(w, "choiceMenu.html", "app/templates/choiceMenu.html", result)
	if err != nil {
		log.Println(fmt.Sprintf("Error in  writing html. %w", err))
	}

	CreateGraph("bar.html", result, choicesOfSorts)
	err = WriteHTML(w, "bar.html", "bar.html", nil)
	if err != nil {
		log.Println(fmt.Sprintf("Error in  writing html. %w", err))
	}

	ctr.Service.CleanService()

}

func (ctr *Controller) GetSorts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := WriteHTML(w, "choiceMenu.html", "app/templates/choiceMenu.html", Times{
		AvailableSorts: ctr.Service.GetAvailableSorts(),
	})

	if err != nil {
		log.Println(fmt.Sprintf("Error in  writing html. %w", err))
	}
}

func CreateGraph(path string, times Times, choicesOfSorts []string) {
	// create a new bar instance
	bar := charts.NewBar()

	// Set global options
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Graph of time dependence on the type of sorting",
		Subtitle: "This is fun to use!",
	}))

	items := make([]opts.BarData, 0)

	for i := 0; i < len(choicesOfSorts); i++ {
		for j := range times.TimesOfSorts {
			if choicesOfSorts[i] == times.TimesOfSorts[j].SortType {
				items = append(items, opts.BarData{Value: times.TimesOfSorts[j].TimeDuration})
				break
			}
		}
	}

	bar.SetXAxis(choicesOfSorts).
		AddSeries("Times", items)
	f, _ := os.Create(path)
	_ = bar.Render(f)
}

func WriteHTML(w http.ResponseWriter, name string, filename string, data interface{}) error {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		//http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return err
	}
	if err = tmpl.Execute(w, data); err != nil {
		//http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return err
	}
	return nil
}
