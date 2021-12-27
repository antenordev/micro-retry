package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"text/template"

	"github.com/hashicorp/go-retryablehttp"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/process", Process)
	http.ListenAndServe(":9090", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/home.html"))

	err := t.Execute(w, Result{})
	if err != nil {
		log.Fatal("Error:", err)
	}

}

func Process(w http.ResponseWriter, r *http.Request) {

	log.Println(r.FormValue("id"))

	result := makeHttpCall("http://localhost:9091", r.FormValue("id"))

	t := template.Must(template.ParseFiles("templates/home.html"))

	err := t.Execute(w, result)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func makeHttpCall(urlMicroService string, id string) Result {

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	values := url.Values{}
	values.Add("id", id)

	res, err := retryablehttp.PostForm(urlMicroService, values)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error:", err)
	}

	result := Result{}
	json.Unmarshal(data, &result)

	return result
}
