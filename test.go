package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	FirstName     string
	LastName      string
	FavoriteFoods []string
	Age           int
}

var p = Person{
	FirstName:     "Harumi",
	LastName:      "Yamashita",
	FavoriteFoods: []string{"Fruits", "Sashimi", "Ramen"},
	Age:           25,
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p)
}
func putHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	fmt.Println(p)

}

func personHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getHandler(w, r)
	case http.MethodPut:
		putHandler(w, r)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/person", personHandler)
	http.ListenAndServe(":8080", nil)
}
