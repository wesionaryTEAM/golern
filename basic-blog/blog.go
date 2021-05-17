package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

/* For Template Caching :: To avoid calling ParseFiles every time a page is rendered */
var templates = template.Must(template.ParseFiles("view.html", "edit.html"))

var validPath = regexp.MustCompile("^/(edit|save|post)/([a-zA-Z0-9]+)$")

type Post struct {
	Title string
	Body  []byte
}

func createAndSavePost(title string, body string) error {
	p := &Post{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		return err
	}
	return nil

}

func (p *Post) save() error {
	fileName := p.Title + ".txt"
	return ioutil.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Post, error) {
	fileName := title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Post{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, templ string, p *Post) {
	err := templates.ExecuteTemplate(w, templ+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		match := validPath.FindStringSubmatch(r.URL.Path)
		if match == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, match[2])
	}
}

func viewPostHandler(w http.ResponseWriter, r *http.Request, title string) {
	body, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", body)
}

func editPostHandler(w http.ResponseWriter, r *http.Request, title string) {
	body, err := loadPage(title)
	if err != nil {
		body = &Post{Title: title}
	}

	renderTemplate(w, "edit", body)
}

func savePostHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	err := createAndSavePost(title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/post/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/post/", makeHandler(viewPostHandler))
	http.HandleFunc("/edit/", makeHandler(editPostHandler))
	http.HandleFunc("/save/", makeHandler(savePostHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
