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
    collection := client.Database("yourdbname").Collection("books")
    controller.SetBookCollection(collection)

    http.HandleFunc("/", controller.BookListHandler)
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}