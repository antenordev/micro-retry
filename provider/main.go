package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":9091", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {

	id := r.PostFormValue("id")

	result := Result{
		Status: "failed",
	}

	if id == "123" {
		result.Status = "success"
	}

	j, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error parsing json", err)
	}

	fmt.Fprintf(w, string(j))
}
