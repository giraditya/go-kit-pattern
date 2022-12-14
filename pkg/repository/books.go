package repository

import (
	"context"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Title     string
	Author    string
	Publisher string
}

type BooksRepository interface {
	Create(ctx context.Context, tx *gorm.DB, books Books) (Books, error)
	Update(ctx context.Context, tx *gorm.DB, books Books) (Books, error)
	Delete(ctx context.Context, tx *gorm.DB, id int) error
	Publish(ctx context.Context, tx *gorm.DB, id int) error
	GetBook(ctx context.Context, tx *gorm.DB, id int) (Books, error)
}

type basicBooksRepository struct{}

func NewBooksRepository() BooksRepository {
	return &basicBooksRepository{}
}

func (repository *basicBooksRepository) Create(ctx context.Context, tx *gorm.DB, books Books) (Books, error) {
	tx.Create(&books)
	return books, tx.Error
}

func (repository *basicBooksRepository) Update(ctx context.Context, tx *gorm.DB, books Books) (Books, error) {
	tx.Updates(Books{Title: books.Title, Author: books.Author, Publisher: books.Publisher})
	return books, tx.Error
}

func (repository *basicBooksRepository) Delete(ctx context.Context, tx *gorm.DB, id int) error {
	tx.Delete(&Books{}, id)
	return tx.Error
}

func (repository *basicBooksRepository) Publish(ctx context.Context, tx *gorm.DB, id int) error {
	tx.Update("publish", 1).Where("id = ?", id)
	return tx.Error
}

func (repository *basicBooksRepository) GetBook(ctx context.Context, tx *gorm.DB, id int) (Books, error) {
	book := Books{}
	row := tx.Table("books").Select("*").Where("id = ?", id).Limit(1).Row()
	if err := row.Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Title, &book.Author, &book.Publisher); err != nil {
		return book, err
	}
	return book, nil
}
