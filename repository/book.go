package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository struct {
	db *pgxpool.Pool
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