package usecase

import "context"

// BookUsecase : BookUsecase is that make book data csv file.
type BookUsecase interface {
	Make(context.Context) error
}
