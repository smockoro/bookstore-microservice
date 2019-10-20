package usecase

import (
	"context"

	"github.com/smockoro/bookstore-microservice/bookshelf/pkg/domain/model"
)

type BookShelfUsecase interface {
	Find(context.Context, string) (*model.Book, error)
	Create(context.Context, *model.Book) (string, error)
}
