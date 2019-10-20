package repository

import (
	"context"

	"github.com/smockoro/bookstore-microservice/bookshelf/pkg/domain/model"
)

type BookShelfRepository interface {
	Create(context.Context, model.Book) (string, error)
	Select(context.Context, string) (model.Book, error)
	Update(context.Context, model.Book) (string, error)
	Delete(context.Context, string) (string, error)
}
