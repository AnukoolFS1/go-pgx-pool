package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"

	"Database/repository"
)

var ConnectionString = "postgres://postgres:bookingstore@localhost:5432/bookstore"

func main() {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping(ctx) // Sends a lightweight query to the database and waits for a response, just to confirm the connection actually works.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgreSQL")

	bookRepo := repository.NewBookRepository(db)

	// bookID, err := bookRepo.CreateBook(
	// 	ctx,
	// 	"The Go Programming Language",
	// 	"Alan Donovan",
	// 	799.00,
	// 	10,
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted book with ID:", bookID)

	book, err := bookRepo.GetBook(ctx, 6)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(book.ID)
	fmt.Println(book.Title)
	fmt.Println(book.Author)
	fmt.Println(book.Price)
	fmt.Println(book.Stock)

	books, err := bookRepo.GetBooks(ctx)

	if err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		fmt.Printf(
			"id: %d, title: %s, author: %s, price: %.2f, stock: %d\n",
			book.ID,
			book.Title,
			book.Author,
			book.Price,
			book.Stock,
		)
	}

}

// "NAMES IDENTIFIER"
// 'String Values'
