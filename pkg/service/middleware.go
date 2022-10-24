package service

import (
	"context"

	log "github.com/go-kit/log"
)

// Middleware describes a service middleware.
type Middleware func(BooksService) BooksService

type loggingMiddleware struct {
	logger log.Logger
	next   BooksService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a BooksService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next BooksService) BooksService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, title string, author string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Create", "title", title, "author", author, "rs", rs, "err", err)
	}()
	return l.next.Create(ctx, title, author)
}
func (l loggingMiddleware) Update(ctx context.Context, title string, author string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Update", "title", title, "author", author, "rs", rs, "err", err)
	}()
	return l.next.Update(ctx, title, author)
}
func (l loggingMiddleware) Delete(ctx context.Context, id int) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "rs", rs, "err", err)
	}()
	return l.next.Delete(ctx, id)
}

func (l loggingMiddleware) Publish(ctx context.Context, id int) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Publish", "id", id, "rs", rs, "err", err)
	}()
	return l.next.Publish(ctx, id)
}
