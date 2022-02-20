package main

import (
	"log"
	"net/http"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to quotebox!"))
}

func (app *application) createBookForm(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/book.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

}

func (app *application) createBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/book", http.StatusSeeOther)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("book_name")
	author := r.PostForm.Get("author_name")
	isbn := r.PostForm.Get("isbn")
	desc := r.PostForm.Get("description")

	insertBook := `
	INSERT INTO books(name, author, isbn, description)
	VALUES($1, $2, $3, $4)`

	_, err = app.db.Exec(insertBook, name, author, isbn, desc)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

}
