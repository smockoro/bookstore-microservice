package repository

import (
	"context"

	"github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/domain/model"
)

// CSVRepository : Make Book data csv file.
type CSVRepository interface {
	Save(context.Context, *model.Book) error
}
