package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

var print = fmt.Printf

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	var filename = p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
	var filename = title + ".txt"
	var body, err = os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return &Page{Title: title, Body: body}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Forda edit ng %s</h1>"+
				"<form action=\"/save/%s\" method=\"POST\">"+
				"<textarea name=\"body\">%s</textarea><br>"+
				"<input type=\"submit\" value=\"Save\">"+
				"</form>",
				p.Title, p.Title, p.Body
		)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foc u guiz, here it iz: %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	var p1 = &Page{Title: "foc u", Body: []byte("Welcome to GEED")}
	p1.save()
	var p2 = loadPage("foc u")
	print(string(p2.Body))
}
