package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Country struct {
	Name string
	Capital string
	Region string
	SubRegion string
	Flag string
	Alpha2Code string
	Alpha3Code string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/api/countries/americas", CountriesHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8087",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func CitiesHandler(writer http.ResponseWriter, request *http.Request) {

}

func CountriesHandler(writer http.ResponseWriter, request *http.Request) {
	_ = json.NewEncoder(writer).Encode(ListCountries());
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))

	if err := t.ExecuteTemplate(writer, "index.html", ListCountries()); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func ListCountries() []Country {
	response, err := http.Get("https://restcountries.eu/rest/v2/region/americas")
	checkError(err)

	data, _ := ioutil.ReadAll(response.Body)
	var countries []Country
	_ = json.Unmarshal([]byte(data), &countries)
	return countries
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}