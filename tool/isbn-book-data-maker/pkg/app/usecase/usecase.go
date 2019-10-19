package usecase

import (
	"context"

	"github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/domain/external/api"
	repo "github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/domain/external/repository"
	Iusecase "github.com/smockoro/bookstore-microservice/tool/isbn-book-data-maker/pkg/domain/usecase"
)

// NewBookUsecase :
func NewBookUsecase(csvRepo repo.CSVRepository, apiClient api.GoogleBookAPIClient) Iusecase.BookUsecase {
	return &bookUsecase{
		csvRepo:   csvRepo,
		apiClient: apiClient,
	}
}

type bookUsecase struct {
	csvRepo   repo.CSVRepository
	apiClient api.GoogleBookAPIClient
}

func (bu *bookUsecase) Make(ctx context.Context) error {

	books := bu.apiClient.Get
	bu.csvRepo.Save(ctx, books)
	return nil
}
