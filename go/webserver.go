package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// what this does is load a page into to var p and have an error if no page
// if there is an error, redirect to edit so it can edit a file (nil means no error)
// we probs dont need it, probs redirect all trafic to home page anyway
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage("homepage")

	if err != nil {
		http.Redirect(w, r, "/home/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

/*
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
*/

var templates = template.Must(template.ParseFiles("view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(home)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/home/", makeHandler(viewHandler))
	//http.HandleFunc("/edit/", makeHandler(editHandler))
	//http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type node struct {
	Course    string  `json:"course"`
	Professor string  `json:"prof"`
	Rating    float32 `json:"rating"`
}

/*
var nodes = []node{
	{Course : "Math", Professor : "jeff", Rating : "3.2"},
	{Course : "Physics", Professor : "saul", Rating : "3.4"},
	{Course : "English", Professor : "white", Rating : "4.6"},
	{Course : "Chemistry", Professor : "jesse", Rating : "4.2"},
	{Course : "History", Professor : "raymond", Rating : "5"},
	{Course : "PE", Professor : "bhat", Rating : "0.0"},
}

*/
/*
for i in range(0,6):
	courses = ["Math", "Physics", "English", "Chemistry", "History", "PE" ]
	professors = ["jeff", "saul", "white", "jesse", "raymond", "bhat"]
	rating = [3.2, 3.4, 4.6, 4.2, 5, 0.0]

	print("\t{{Course : \"{0}\", Professor : \"{1}\", Rating : \"{2}\"}},".format(courses[i], professors[i], rating[i]))
*/
