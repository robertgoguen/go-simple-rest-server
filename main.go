package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Key struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type allKey []Key

var keys = allKey{
	{
		Key:   "quincy",
		Value: "rules",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createKey(w http.ResponseWriter, r *http.Request) {
	var newKey Key
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter value with the key value only in order to update")
	}

	json.Unmarshal(reqBody, &newKey)
	keys = append(keys, newKey)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newKey)
}

func getOneKey(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]

	for _, singleKey := range keys {
		if singleKey.Key == key {
			json.NewEncoder(w).Encode(singleKey)
		}
	}
}

func getAllKeys(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(keys)
}

func updateKey(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]
	var updatedKey Key

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter value with the key value only in order to update")
	}
	json.Unmarshal(reqBody, &updatedKey)

	for i, singleKey := range keys {
		if singleKey.Key == key {
			singleKey.Value = updatedKey.Value
			keys = append(keys[:i], singleKey)
			json.NewEncoder(w).Encode(singleKey)
		}
	}
}

func deleteKey(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]

	for i, singleKey := range keys {
		if singleKey.Key == key {
			keys = append(keys[:i], keys[i+1:]...)
			fmt.Fprintf(w, "The key with Key %v has been deleted successfully\n", key)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/key", createKey).Methods("POST")
	router.HandleFunc("/keys", getAllKeys).Methods("GET")
	router.HandleFunc("/keys/{id}", getOneKey).Methods("GET")
	router.HandleFunc("/keys/{id}", updateKey).Methods("PATCH")
	router.HandleFunc("/keys/{id}", deleteKey).Methods("DELETE")
	fmt.Printf("ListenAndServe at address :8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
