package main

import (
	"fmt"
	"net/http"
)

var squeaker = NewSqueaker()

const lenPath = len("/s/")

func main() {
	squeaker.Add("hello", "World! This is a test.")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/s/", topicHandler)
	http.HandleFunc("/comment/", commentHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Home</h1>")
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
	fmt.Fprintf(w, "<h1>%s</h1>", title)
	for _, v := range squeaker.topics[title] {
		fmt.Fprintf(w, "<div><h2>%s</h2><p>%s</p><p>%v</p></div>", v.UUID, v.Message, v.Time)
	}
}

func commentHandler(w http.ResponseWriter, r *http.Request) {
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
}
