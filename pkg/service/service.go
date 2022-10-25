package service

import (
	"books/pkg/repository"
	"context"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

// BooksService describes the service.
type BooksService interface {
	// Add your methods here
	Create(ctx context.Context, title string, author string) (rs string, err error)
	Update(ctx context.Context, title string, author string) (rs string, err error)
	Delete(ctx context.Context, id int) (rs string, err error)
	Publish(ctx context.Context, id int) (rs string, err error)
	GetBook(ctx context.Context, id int) (rs repository.Books, err error)
	SendEmailBookPublished(ctx context.Context, id int) (rs string, err error)
}

type basicBooksService struct {
	DB              *gorm.DB
	Redis           *redis.Client
	BooksRepository repository.BooksRepository
}

func (b *basicBooksService) Create(ctx context.Context, title string, author string) (rs string, err error) {
	// TODO implement the business logic of Create
	books := repository.Books{
		Title:  title,
		Author: author,
	}
	tx := b.DB.Begin()
	_, err = b.BooksRepository.Create(ctx, tx, books)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	rs = "books created"
	return rs, err
}
func (b *basicBooksService) Update(ctx context.Context, title string, author string) (rs string, err error) {
	// TODO implement the business logic of Update
	books := repository.Books{
		Title:  title,
		Author: author,
	}
	tx := b.DB.Begin()
	_, err = b.BooksRepository.Update(ctx, tx, books)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	rs = "books updated"
	return rs, err
}
func (b *basicBooksService) Delete(ctx context.Context, id int) (rs string, err error) {
	// TODO implement the business logic of Delete
	tx := b.DB.Begin()
	err = b.BooksRepository.Delete(ctx, tx, id)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	rs = "books created"
	return rs, err
}
func (b *basicBooksService) Publish(ctx context.Context, id int) (rs string, err error) {
	// TODO implement the business logic of Publish
	tx := b.DB.Begin()
	err = b.BooksRepository.Publish(ctx, tx, id)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	rs = "books published"
	return rs, err
}
func (b *basicBooksService) GetBook(ctx context.Context, id int) (rs repository.Books, err error) {
	// TODO implement the business logic of GetBook
	rs, err = b.BooksRepository.GetBook(ctx, b.DB, id)
	return rs, err
}
func (b *basicBooksService) SendEmailBookPublished(ctx context.Context, id int) (rs string, err error) {
	// TODO implement the business logic of SendEmailBookPublished

	return rs, err
}

// NewBasicBooksService returns a naive, stateless implementation of BooksService.
func NewBasicBooksService(db *gorm.DB, redis *redis.Client, bookRepository repository.BooksRepository) BooksService {
	return &basicBooksService{
		DB:              db,
		Redis:           redis,
		BooksRepository: bookRepository,
	}
}

// New returns a BooksService with all of the expected middleware wired in.
func New(middleware []Middleware, db *gorm.DB, redis *redis.Client, bookRepository repository.BooksRepository) BooksService {
	var svc BooksService = NewBasicBooksService(db, redis, bookRepository)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
