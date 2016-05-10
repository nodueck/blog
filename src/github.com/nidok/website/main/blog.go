package main

import (

	"net/http"
	// "github.com/russross/blackfriday"
	"html/template"
	"fmt"
)


type AboutMe struct {
	name string
	title string
	shortText string
}

type SiteMetaInformation struct {
	title string
	useBootstrap bool
	usejQuery bool
}

func main() {
	/*http.HandleFunc("/markdown", GenerateMarkDown)
	http.Handle("/", http.FileServer(http.Dir("../html")))
	http.ListenAndServe(":8080", nil)*/

	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	me := AboutMe{"Nikolai", "Developer", "It'me :'D"}
	// meta := SiteMetaInformation{"Blog-Test", true, true}

	t := template.Must(template.New("Home").ParseFiles("../html/index.html"))
	Println(t)
	t.Execute(w, me)
}