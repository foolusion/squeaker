package main

import (
	"fmt"
	"net/http"
)

var squeaker = NewSqueaker()

const (
		lenPath = len("/s/")
		lenSqueak = len("/squeak/")
		lenSave = len("/save/")
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
	for k, v := range squeaker.topics {
		if count < 100 {
			fmt.Fprintf(w, "<div><a href=\"/s/%s\">%s</a> %d squeaks</div>", k, k, len(v))
		} else {
			break
		}
	}
}

func topicHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	fmt.Fprintf(w, "<h1>%s</h1><a href=\"/squeak/%s\">squeak</a><a href=\"/\">Home</a>", title, title)
	for _, v := range squeaker.topics[title] {
		fmt.Fprintf(w, "<div><h2>%s</h2><p>%s</p><p>%v</p></div>", v.Message, v.UUID, v.Time)
	}
}

func squeakHandler(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[lenSqueak:]
		fmt.Fprintf(w,
		"<h1> Squeaking on %s</h1>" +
		"<form action=\"/save/%s\" method=\"POST\">" +
		"<textarea name=\"message\" rows=\"2\" cols=\"80\" maxlength=\"140\"></textarea>" +
		"<input type=\"submit\" value=\"squeak\">" +
		"</form>",
		title, title)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[lenSave:]
		message := r.FormValue("message")
		squeaker.Add(title, message)
		http.Redirect(w, r, "/s/"+title, http.StatusFound)
}
