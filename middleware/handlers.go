package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
type Stock struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Count int64  `json:"count,omitempty"`
}

func GetStockById(w http.ResponseWriter, r *http.Request) {
	log.Println("TEst")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert id into number")
	}

	respone := Stock{
		ID:    int64(id),
		Name:  "Test",
		Count: 23,
	}

	json.NewEncoder(w).Encode(respone)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode value %v", err)
	}

	//Insert in db and return id from db

	response := Response{
		ID:      23,
		Message: "Added new stock",
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var stock Stock

	stockId, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to decode value %v", err)
	}

	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to read body %v", err)
	}

	log.Printf("%v", stock)
	//Insert in db and return id from db

	response := Response{
		ID:      int64(stockId),
		Message: "Stack updated",
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)

	stockId, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to read params %v", err)
	}

	//Add module for delete Stock

	response := Response{
		ID:      int64(stockId),
		Message: "Stack deleted",
	}

	json.NewEncoder(w).Encode(response)
}
