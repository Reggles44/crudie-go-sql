package main

import (
	"encoding/json"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {return}
	
	var data GoFooBar
	err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {return}
	
	result := db.Create(&data)
	if result.Error != nil{
		panic(result.Error)
	}
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(data)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {return}
	
	foo := r.FormValue("foo")
	
	var data GoFooBar
	result := db.Model(GoFooBar{Foo: foo}).First(&data)
	if result.Error != nil{
		panic(result.Error)
	}
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {return}
	
	var requestData GoFooBar
	err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {return}
	
	var data GoFooBar
	result := db.Model(GoFooBar{Foo: requestData.foo}).First(&data)
	if result.Error != nil{
		panic(result.Error)
	}
	
	data.bar = requestData.bar
	result := db.Save(&data)
	if result.Error != nil{
		panic(result.Error)
	}
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {return}
	
	var data GoFooBar
	result := db.Model(GoFooBar{Foo: foo}).First(&data)
	if result.Error != nil{
		panic(result.Error)
	}
	
	result := db.Delete(&data)
	if result.Error != nil{
		panic(result.Error)
	}
	
}