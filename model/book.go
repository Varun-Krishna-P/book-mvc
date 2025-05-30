package model

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "log"
    "time"
)

type Book struct {
    ID     int    `bson:"id"`
    Title  string `bson:"title"`
    Author string `bson:"author"`
}

// GetAllBooks fetches all books from the given MongoDB collection
func GetAllBooks(collection *mongo.Collection) []Book {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.D{})
    if err != nil {
        log.Println("Error fetching books:", err)
        return nil
    }
    defer cursor.Close(ctx)

    var books []Book
    for cursor.Next(ctx) {
        var book Book
        if err := cursor.Decode(&book); err != nil {
            log.Println("Error decoding book:", err)
            continue
        }
        books = append(books, book)
    }
    if err := cursor.Err(); err != nil {
        log.Println("Cursor error:", err)
    }
    return books
}
