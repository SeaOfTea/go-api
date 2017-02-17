package main

import (
	"log"
	"net/http"
	"time"

	"fmt"
	"html"

	"encoding/json"

	"github.com/gorilla/mux"
)

type todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type todos []todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/todos", todoIndex)
	router.HandleFunc("/todos/{todoId}", todoShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	innerTodos := todos{
		todo{Name: "Write presentation"},
		todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(innerTodos)
}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
