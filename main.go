package main

import (
	"fmt"
	"net/http"

	"github.com/foolusion/squeaker/squeaker"
)

var sq = squeaker.NewMapSqueaker()

const (
	lenPath   = len("/s/")
	lenSqueak = len("/squeak/")
	lenSave   = len("/save/")
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/s/", topicHandler)
	http.HandleFunc("/squeak/", squeakHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Squeaker</h1>")
	count := 0
	for _, v := range sq.Topics() {
		if count < 100 {
			fmt.Fprintf(w, "<div><a href=\"/s/%s\">%s</a> %d squeaks</div>", v,
				v, len(sq.Get(v)))
		} else {
			break
		}
	}
}

func topicHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	squeaks := sq.Get(title)
	fmt.Fprintf(w, "<h1>%s</h1><a href=\"/\">Home</a> <a href=\"/squeak/%s\">squeak</a>", title, title)
	for _, v := range squeaks {
		fmt.Fprintf(w, "<div><h2>%s</h2><p>%s</p><p>%v</p></div>", v.Message, v.UUID, v.Time)
	}
	if len(squeaks) == 0 {
		fmt.Fprint(w, "<div><h2>No Squeaks Yet</h2></div>")
	}
}

func squeakHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenSqueak:]
	fmt.Fprintf(w,
		"<h1> Squeaking on %s</h1>"+
			"<form action=\"/save/%s\" method=\"POST\">"+
			"<textarea name=\"message\" rows=\"2\" cols=\"80\" maxlength=\"140\"></textarea>"+
			"<input type=\"submit\" value=\"squeak\">"+
			"</form>",
		title, title)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenSave:]
	message := r.FormValue("message")
	sq.Squeak(title, message)
	http.Redirect(w, r, "/s/"+title, http.StatusFound)
}
