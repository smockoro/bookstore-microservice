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
	MockList := []string{
		"9784297107277",
		"9784295000969",
		"9784798142418",
		"9784908686030",
		"9784873118468",
		"9784873117522",
		"9784774183923",
	}
	for _, isbn := range MockList {
		book, err := bu.apiClient.Get(isbn)
		if err != nil {
			return err
		}
		if book != nil {
			err = bu.csvRepo.Save(ctx, book)
			if err != nil {
				return err
			}
		}

	}
	/*
		for i := 800000000; i <= 999999999; i++ {
			isbn := fmt.Sprintf("%d%d%d", model.PREFIX, model.JA, i)
			fmt.Println("isbn: ", isbn)
			book, err := bu.apiClient.Get(isbn)
			if err != nil {
				return err
			}
			if book != nil {
				err = bu.csvRepo.Save(ctx, book)
				if err != nil {
					return err
				}
			}
		}
	*/
	return nil
}
