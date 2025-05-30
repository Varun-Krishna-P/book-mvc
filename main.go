package main

import (
    "net/http"
    "book-mvc/controller"
    "book-mvc/config"
    "context"
    "fmt"
)

func main() {
    // Connect to MongoDB
***REMOVED***
    if err != nil {
        fmt.Println("MongoDB connection error:", err)
        return
    }
    defer client.Disconnect(context.Background())

    // Get the books collection from your database (replace "yourdbname" and "books" as needed)
    books_collection := client.Database("learning_db").Collection("books")
    controller.SetBookCollection(books_collection)

    http.HandleFunc("/", controller.BookListHandler)
    http.HandleFunc("/books/new", controller.NewBookFormHandler)
    http.HandleFunc("/books/create", controller.CreateBookHandler)
    http.HandleFunc("/books/show", controller.ShowBookHandler)
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}