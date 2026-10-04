package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository struct {
	db *pgxpool.Pool
}

type Book struct {
    ID     int
    Title  string
    Author string
    Price  float32
    Stock  int
}

func NewBookRepository(db *pgxpool.Pool) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) CreateBook(
	ctx context.Context,
	title string,
	author string,
	price float64,
	stock int,
) (int, error) {

	var bookID int

	err := r.db.QueryRow(
		ctx,
		`INSERT INTO books (title, author, price, stock)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		title,
		author,
		price,
		stock,
	).Scan(&bookID)

	if err != nil {
		return 0, err
	}

	return bookID, nil
}

func (r *BookRepository) GetBook(ctx context.Context, id int) (Book, error) {
    var book Book

    err := r.db.QueryRow(
        ctx,
        `SELECT id, title, author, price, stock
         FROM books
         WHERE id = $1`,
        id,
    ).Scan(
        &book.ID,
        &book.Title,
        &book.Author,
        &book.Price,
        &book.Stock,
    )

    if err != nil {
        return Book{}, err
    }

    return book, nil
}