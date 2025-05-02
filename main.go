package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func postTaskHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task = body.Task
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "task set to: %s", task)
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", task)
}

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

func main() {
	http.HandleFunc("/get", getTaskHandler)
	http.HandleFunc("/post", postTaskHandler)
	fmt.Println("Server start http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
