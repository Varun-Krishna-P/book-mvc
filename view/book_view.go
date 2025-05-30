package view

import (
    "fmt"
    "net/http"
    "book-mvc/model"
)

func RenderBookList(w http.ResponseWriter, books []model.Book) {
    fmt.Fprint(w, "<h1>Book List</h1><ul>")
    for _, book := range books {
        fmt.Fprintf(w, "<li>%s by %s</li>", book.Title, book.Author)
    }
    fmt.Fprint(w, "</ul>")
}
