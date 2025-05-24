package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := "Hello, World!"
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatal(err)
	}

}
func Greetings(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	w.Header().Set("Content-Type", "application/json")
	note := "\nHow are you doing today?\nJust wanted to remind you that you are amazing\nYou are destined for greatness!"
	fmt.Fprint(w, "Hello "+name, note)
}
