package controller

import (
	"encoding/json"
	"errors"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"html/template"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"sync"
)

type Controller struct {
	Service service
	sync.RWMutex
}

type Times struct {
	TimesOfSorts   []TimeOfSort
	AvailableSorts []string
	StartedArray   []int
	SortedArray    []int
}

type TimeOfSort struct {
	SortType     string  `json:"type"`
	TimeDuration float64 `json:"time"`
}

type Error struct {
	Err    string
	Status int
}

type service interface {
	StartSorting([]string) error
	FillByRand(n int)
	FillFromFile(path string) error
	SetArrayByUserChoice(interface{}, []string) error
	GetSortsResultJSON() (string, error)
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var input interface{}

	if filename := r.PostFormValue("filenameArr"); filename != "" { // user input is filename
		input = filename
	} else if sizeRandArr := r.PostFormValue("sizeRandArr"); sizeRandArr != "" { // user input is array length
		size, err := strconv.Atoi(sizeRandArr) // string to int
		if err != nil {
			WriteError(w, "app/templates/errorMenu.html", http.StatusBadRequest, err) // write error
			return
		}
		input = size
	}

	choicesOfSorts := r.PostForm["checkbox"] // check which checkbox chose
	if len(choicesOfSorts) == 0 {
		WriteError(w, "app/templates/errorMenu.html", http.StatusBadRequest, errors.New("You need to choose one or more checkbox"))
		return
	}

	err := ctr.Service.SetArrayByUserChoice(input, choicesOfSorts) // fill aray by choice
	if err != nil {
		WriteError(w, "app/templates/errorMenu.html", http.StatusBadRequest, err)
		return
	}

	dataJSON, err := ctr.Service.GetSortsResultJSON() // get json info about time of all sorts
	if err != nil {
		WriteError(w, "app/templates/errorMenu.html", http.StatusBadRequest, err)
		return
	}

	var dataStruct map[string][]TimeOfSort
	err = json.Unmarshal([]byte(dataJSON), &dataStruct) // decode from json to map
	if err != nil {
		WriteError(w, "app/templates/errorMenu.html", http.StatusBadRequest, err)
		return
	}

	result := Times{ // create result struct for html
		StartedArray:   ctr.Service.GetStartedArray(),
		AvailableSorts: ctr.Service.GetAvailableSorts(),
		SortedArray:    ctr.Service.GetSortedArray(),
		TimesOfSorts:   dataStruct["Sorts"],
	}

	err = WriteHTML(w, "app/templates/choiceMenu.html", result)
	if err != nil {
		WriteError(w, "app/templates/errorMenu.html", http.StatusInternalServerError, err)
		return
	}

	err = CreateGraph("bar.html", result, choicesOfSorts) // create graphic
	if err != nil {
		WriteError(w, "app/templates/errorMenu.html", http.StatusInternalServerError, err)
		return
	}

	err = WriteHTML(w, "bar.html", nil) //write graphic
	if err != nil {
		WriteError(w, "app/templates/errorMenu.html", http.StatusInternalServerError, err)
		return
	}
}

func (ctr *Controller) GetSorts(w http.ResponseWriter, r *http.Request) { //basic menu
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := WriteHTML(w, "app/templates/choiceMenu.html", Times{
		AvailableSorts: ctr.Service.GetAvailableSorts(),
	})

	if err != nil {
		WriteError(w, "app/templates/errorMenu.html", http.StatusInternalServerError, err)
		return
	}
}

func CreateGraph(path string, times Times, choicesOfSorts []string) error {
	bar := charts.NewBar()

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
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	err = bar.Render(f)
	if err != nil {
		return err
	}
	return nil
}

func WriteHTML(w http.ResponseWriter, filename string, data interface{}) error {
	tmpl, err := template.ParseFiles(filename) // found html file
	if err != nil {
		//http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return err
	}
	if err = tmpl.Execute(w, data); err != nil {
		//http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return err
	} // execute with interface
	return nil
}

func WriteError(w http.ResponseWriter, filename string, statusCode int, errTitle error) {
	err := WriteHTML(w, filename, Error{
		Status: statusCode,
		Err:    errTitle.Error(),
	}) // write error menu
	if err != nil {
		log.Fatalln(err)
		return
	}
}
