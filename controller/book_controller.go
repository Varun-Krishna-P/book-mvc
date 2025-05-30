package controller

import (
    "net/http"
    "book-mvc/model"
    "go.mongodb.org/mongo-driver/mongo"
    "html/template"
    "log"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var bookCollection *mongo.Collection

// SetBookCollection allows main.go to inject the MongoDB collection
func SetBookCollection(collection *mongo.Collection) {
    bookCollection = collection
}

func BookListHandler(w http.ResponseWriter, r *http.Request) {
    if bookCollection == nil {
        http.Error(w, "Database not initialized", http.StatusInternalServerError)
        return
    }
    books := model.GetAllBooks(bookCollection)
    tmpl := template.Must(template.ParseFiles("view/books.html"))
    tmpl.Execute(w, books)
}

// Show the form to add a new book
func NewBookFormHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("view/new_book.html"))
    tmpl.Execute(w, nil)
}

// Handle the POST request to create a new book
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        http.Redirect(w, r, "/books/new", http.StatusSeeOther)
        return
    }

    title := r.FormValue("title")
    author := r.FormValue("author")

    book := model.Book{
        Title:  title,
        Author: author,
    }

    if err := model.InsertBook(bookCollection, &book); err != nil {
        log.Println("Error inserting book:", err)
        http.Error(w, "Failed to add book", http.StatusInternalServerError)
        return
    }else {
		log.Println("Book added successfully:", book)
	}
	// Redirect to the book list after successful insertion	

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ShowBookHandler(w http.ResponseWriter, r *http.Request) {
    if bookCollection == nil {
        http.Error(w, "Database not initialized", http.StatusInternalServerError)
        return
    }
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "Missing book ID", http.StatusBadRequest)
        return
    }
    objID, err := primitive.ObjectIDFromHex(idStr)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
    book, err := model.GetBookByID(bookCollection, objID)
    if err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    tmpl := template.Must(template.ParseFiles("view/book.html"))
    tmpl.Execute(w, book)
}