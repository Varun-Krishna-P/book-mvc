package controller

import (
    "net/http"
    "book-mvc/model"
    "go.mongodb.org/mongo-driver/mongo"
    "html/template"
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