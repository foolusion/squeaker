package main

import (
	"html/template"
	"net/http"

	"github.com/foolusion/squeaker/squeak"
)

var sq = squeak.NewMapSqueaker()

const (
	lenPath   = len("/s/")
	lenSqueak = len("/squeak/")
	lenSave   = len("/save/")
)

type page struct {
	Title   string
	Squeaks []squeak.Squeak
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/s/", topicHandler)
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/squeak/", squeakHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, sq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func topicHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	squeaks := sq.Get(title)
	p := page{title, squeaks}
	t, err := template.ParseFiles("view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func squeakHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenSqueak:]
	p := page{Title: title}
	t, err := template.ParseFiles("squeak.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenSave:]
	message := r.FormValue("message")
	sq.Squeak(title, message)
	http.Redirect(w, r, "/s/"+title, http.StatusFound)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("q")
	http.Redirect(w, r, "/s/"+title, http.StatusFound)
}
